package api_test

import (
	"context"
	"database/sql"

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

		BeforeEach(func() {
			createRequest = &desc.CreateNoteV1Request{
				UserId:      1,
				ClassroomId: 1,
				DocumentId:  1,
			}

			rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

			mock.ExpectQuery("INSERT INTO notes").
				WithArgs(createRequest.UserId, createRequest.ClassroomId, createRequest.DocumentId).
				WillReturnRows(rows)

			createResponse, err = grpcApi.CreateNoteV1(ctx, createRequest)
		})

		It("good creating", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.NoteId).Should(Equal(uint64(1)))
		})
	})
})
