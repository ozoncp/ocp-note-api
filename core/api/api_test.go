package api_test

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-note-api/core/api"
	"github.com/ozoncp/ocp-note-api/core/repo"
	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
)

var _ = Describe("Api", func() {

	var (
		ctx context.Context

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		// note = []note.Note{
		// 	{Id: 1, UserId: 1, ClassroomId: 1, DocumentId: 1},
		// 	{Id: 2, UserId: 2, ClassroomId: 2, DocumentId: 2},
		// }

		storage repo.Repo
		grpcApi desc.OcpNoteApiServer

		createRequest  *desc.CreateNoteV1Request
		createResponse *desc.CreateNoteV1Response

		err error
	)

	BeforeEach(func() {
		ctx = context.Background()

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")

		storage = repo.New(*sqlxDB)
		grpcApi = api.NewOcpNoteApi(storage)
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("create note", func() {

		var id uint64 = 1

		BeforeEach(func() {
			createRequest = &desc.CreateNoteV1Request{
				UserId:      1,
				ClassroomId: 1,
				DocumentId:  1,
			}

			mock.ExpectQuery("INSERT INTO notes").
				WithArgs(createRequest.UserId, createRequest.ClassroomId, createRequest.DocumentId).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

			createResponse, err = grpcApi.CreateNoteV1(ctx, createRequest)
		})

		It("successful creation of a note in the database", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.NoteId).Should(Equal(id))
		})
	})

	Context("create note with invalid arguments", func() {

		BeforeEach(func() {
			createRequest = &desc.CreateNoteV1Request{
				UserId:      -1,
				ClassroomId: 1,
				DocumentId:  1,
			}

			// setting the wait for the mock request is not required,
			// since the error will return earlier due to invalid arguments

			createResponse, err = grpcApi.CreateNoteV1(ctx, createRequest)
		})

		It("failed note creation due to invalid arguments", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("unsuccessful note creation", func() {

		BeforeEach(func() {
			createRequest = &desc.CreateNoteV1Request{
				UserId:      1,
				ClassroomId: 1,
				DocumentId:  1,
			}

			mock.ExpectQuery("INSERT INTO notes").
				WithArgs(createRequest.UserId, createRequest.ClassroomId, createRequest.DocumentId).
				WillReturnError(errors.New("failed to execute sql request"))

			createResponse, err = grpcApi.CreateNoteV1(ctx, createRequest)
		})

		It("failed to execute sql request", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})
})
