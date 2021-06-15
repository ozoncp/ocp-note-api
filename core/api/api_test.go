package api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-note-api/core/api"
	"github.com/ozoncp/ocp-note-api/core/mocks"
	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	desc "github.com/ozoncp/ocp-note-api/pkg/ocp-note-api"
)

var _ = Describe("Api", func() {

	var (
		ctx  context.Context
		ctrl *gomock.Controller

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		notes = []note.Note{
			{Id: 1, UserId: 1, ClassroomId: 1, DocumentId: 1},
			{Id: 2, UserId: 2, ClassroomId: 2, DocumentId: 2},
		}

		storage          repo.Repo
		dataProducerMock *mocks.MockProducer
		grpcApi          desc.OcpNoteApiServer

		createRequest  *desc.CreateNoteV1Request
		createResponse *desc.CreateNoteV1Response

		multiCreateNotesV1Request  *desc.MultiCreateNotesV1Request
		multiCreateNotesV1Response *desc.MultiCreateNotesV1Response

		updateNoteV1Request  *desc.UpdateNoteV1Request
		updateNoteV1Response *desc.UpdateNoteV1Response

		describeRequest  *desc.DescribeNoteV1Request
		describeResponse *desc.DescribeNoteV1Response

		listNotesV1Request  *desc.ListNotesV1Request
		listNotesV1Response *desc.ListNotesV1Response

		removeNoteV1Request  *desc.RemoveNoteV1Request
		removeNoteV1Response *desc.RemoveNoteV1Response

		err       error
		chunkSize uint32
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")

		chunkSize = 5
		storage = repo.New(*sqlxDB, chunkSize)
		dataProducerMock = mocks.NewMockProducer(ctrl)

		grpcApi = api.NewOcpNoteApi(storage, dataProducerMock)
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
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

			dataProducerMock.EXPECT().Send(gomock.Any())

			createResponse, err = grpcApi.CreateNoteV1(ctx, createRequest)
		})

		It("successful creation of a note in the database", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.NoteId).Should(Equal(id))
		})
	})

	Context("multi create notes with invalid arguments", func() {

		BeforeEach(func() {
			multiCreateNotesV1Request = &desc.MultiCreateNotesV1Request{
				Notes: []*desc.NewNote{{
					UserId:      -1,
					ClassroomId: 1,
					DocumentId:  1,
				}},
			}

			// setting the wait for the mock request is not required,
			// since the error will return earlier due to invalid arguments

			multiCreateNotesV1Response, err = grpcApi.MultiCreateNotesV1(ctx, multiCreateNotesV1Request)
		})

		It("failed notes multi creation due to invalid arguments", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateNotesV1Response).Should(BeNil())
		})
	})

	Context("unsuccessful notes multi creation", func() {

		BeforeEach(func() {
			multiCreateNotesV1Request = &desc.MultiCreateNotesV1Request{
				Notes: []*desc.NewNote{
					{
						UserId:      int32(notes[0].UserId),
						ClassroomId: int32(notes[0].ClassroomId),
						DocumentId:  int32(notes[0].DocumentId),
					},
					{
						UserId:      int32(notes[1].UserId),
						ClassroomId: int32(notes[1].ClassroomId),
						DocumentId:  int32(notes[1].DocumentId),
					}},
			}

			mock.ExpectExec("INSERT INTO notes").
				WithArgs(notes[0].UserId, notes[0].ClassroomId, notes[0].DocumentId, notes[1].UserId, notes[1].ClassroomId, notes[1].DocumentId).
				WillReturnError(errors.New("failed to execute sql request"))

			multiCreateNotesV1Response, err = grpcApi.MultiCreateNotesV1(ctx, multiCreateNotesV1Request)
		})

		It("failed to execute sql request", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateNotesV1Response).Should(BeNil())
		})
	})

	Context("multi create notes", func() {

		BeforeEach(func() {
			multiCreateNotesV1Request = &desc.MultiCreateNotesV1Request{
				Notes: []*desc.NewNote{
					{
						UserId:      int32(notes[0].UserId),
						ClassroomId: int32(notes[0].ClassroomId),
						DocumentId:  int32(notes[0].DocumentId),
					},
					{
						UserId:      int32(notes[1].UserId),
						ClassroomId: int32(notes[1].ClassroomId),
						DocumentId:  int32(notes[1].DocumentId),
					}},
			}

			mock.ExpectExec("INSERT INTO notes").
				WithArgs(notes[0].UserId, notes[0].ClassroomId, notes[0].DocumentId, notes[1].UserId, notes[1].ClassroomId, notes[1].DocumentId).
				WillReturnResult(sqlmock.NewResult(0, 2))

			multiCreateNotesV1Response, err = grpcApi.MultiCreateNotesV1(ctx, multiCreateNotesV1Request)
		})

		It("successful multi creation of a notes in the database", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateNotesV1Response.NumberOfNotesCreated).Should(Equal(uint64(len(notes))))
		})
	})

	Context("update note with invalid arguments", func() {

		BeforeEach(func() {
			updateNoteV1Request = &desc.UpdateNoteV1Request{
				Note: &desc.Note{
					Id:          1,
					UserId:      -1,
					ClassroomId: 10,
					DocumentId:  20,
				},
			}

			// setting the wait for the mock request is not required,
			// since the error will return earlier due to invalid arguments

			updateNoteV1Response, err = grpcApi.UpdateNoteV1(ctx, updateNoteV1Request)
		})

		It("failed note updating due to invalid arguments", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(updateNoteV1Response).Should(BeNil())
		})
	})

	Context("unsuccessful note updating", func() {

		BeforeEach(func() {
			updateNoteV1Request = &desc.UpdateNoteV1Request{
				Note: &desc.Note{
					Id:          1,
					UserId:      1,
					ClassroomId: 10,
					DocumentId:  20,
				},
			}

			mock.ExpectExec("UPDATE notes").
				WithArgs(updateNoteV1Request.Note.UserId, updateNoteV1Request.Note.ClassroomId, updateNoteV1Request.Note.DocumentId, updateNoteV1Request.Note.Id).
				WillReturnError(errors.New("failed to execute sql request"))

			updateNoteV1Response, err = grpcApi.UpdateNoteV1(ctx, updateNoteV1Request)
		})

		It("failed to execute sql request", func() {
			Expect(updateNoteV1Response.Found).Should(Equal(false))
		})
	})

	Context("update note", func() {

		BeforeEach(func() {
			updateNoteV1Request = &desc.UpdateNoteV1Request{
				Note: &desc.Note{
					Id:          1,
					UserId:      1,
					ClassroomId: 10,
					DocumentId:  20,
				},
			}

			mock.ExpectExec("UPDATE notes").
				WithArgs(updateNoteV1Request.Note.UserId, updateNoteV1Request.Note.ClassroomId, updateNoteV1Request.Note.DocumentId, updateNoteV1Request.Note.Id).
				WillReturnResult(sqlmock.NewResult(0, 1))

			dataProducerMock.EXPECT().Send(gomock.Any())

			updateNoteV1Response, err = grpcApi.UpdateNoteV1(ctx, updateNoteV1Request)
		})

		It("successful updating of a note in the database", func() {
			Expect(err).Should(BeNil())
			Expect(updateNoteV1Response.Found).Should(Equal(true))
		})
	})

	Context("describe note with invalid arguments", func() {

		var id int64 = 1

		BeforeEach(func() {
			describeRequest = &desc.DescribeNoteV1Request{
				NoteId: id,
			}

			// setting the wait for the mock request is not required,
			// since the error will return earlier due to invalid arguments

			describeResponse, err = grpcApi.DescribeNoteV1(ctx, describeRequest)
		})

		It("could not get the description of the note due to invalid arguments", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(describeResponse).Should(BeNil())
		})
	})

	Context("unsuccessful receipt of the description of the note", func() {

		var id int64 = 1

		BeforeEach(func() {
			describeRequest = &desc.DescribeNoteV1Request{
				NoteId: id,
			}

			mock.ExpectQuery("SELECT (.+) FROM notes WHERE").
				WithArgs(describeRequest.NoteId).
				WillReturnError(errors.New("failed to execute sql request"))

			describeResponse, err = grpcApi.DescribeNoteV1(ctx, describeRequest)
		})

		It("failed to execute sql request", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(describeResponse).Should(BeNil())
		})
	})

	Context("describe note", func() {

		var (
			id           uint64 = 1
			user_id      uint32 = 1
			classroom_id uint32 = 1
			document_id  uint32 = 1
		)

		BeforeEach(func() {
			describeRequest = &desc.DescribeNoteV1Request{
				NoteId: int64(id),
			}

			mock.ExpectQuery("SELECT (.+) FROM notes WHERE").
				WithArgs(describeRequest.NoteId).
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "user_id", "classroom_id", "document_id"}).
					AddRow(id, user_id, classroom_id, document_id))

			describeResponse, err = grpcApi.DescribeNoteV1(ctx, describeRequest)
		})

		It("successful receipt of the note description", func() {
			Expect(err).Should(BeNil())
			Expect(describeResponse.Note.Id).Should(Equal(int64(id)))
			Expect(describeResponse.Note.UserId).Should(Equal(int32(user_id)))
			Expect(describeResponse.Note.ClassroomId).Should(Equal(int32(classroom_id)))
			Expect(describeResponse.Note.DocumentId).Should(Equal(int32(document_id)))
		})
	})

	Context("list notes with invalid arguments", func() {

		var (
			limit  uint64 = 0
			offset uint64 = 0
		)

		BeforeEach(func() {
			listNotesV1Request = &desc.ListNotesV1Request{
				Limit:  int64(limit),
				Offset: int64(offset),
			}

			// setting the wait for the mock request is not required,
			// since the error will return earlier due to invalid arguments

			listNotesV1Response, err = grpcApi.ListNotesV1(ctx, listNotesV1Request)
		})

		It("failed to get the list of notes due to invalid arguments", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(listNotesV1Response).Should(BeNil())
		})
	})

	Context("unsuccessful retrieval of the list of notes", func() {

		var (
			limit  uint64 = 10
			offset uint64 = 1
		)

		BeforeEach(func() {
			listNotesV1Request = &desc.ListNotesV1Request{
				Limit:  int64(limit),
				Offset: int64(offset),
			}

			query := fmt.Sprintf("SELECT (.+) FROM notes LIMIT %d OFFSET %d", listNotesV1Request.Limit, listNotesV1Request.Offset)
			mock.ExpectQuery(query).
				WillReturnError(errors.New("failed to execute sql request"))

			listNotesV1Response, err = grpcApi.ListNotesV1(ctx, listNotesV1Request)
		})

		It("failed to execute sql request", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(listNotesV1Response).Should(BeNil())
		})
	})

	Context("list notes", func() {

		var (
			limit  uint64 = 10
			offset uint64 = 1
		)

		BeforeEach(func() {
			listNotesV1Request = &desc.ListNotesV1Request{
				Limit:  int64(limit),
				Offset: int64(offset),
			}

			query := fmt.Sprintf("SELECT (.+) FROM notes LIMIT %d OFFSET %d", listNotesV1Request.Limit, listNotesV1Request.Offset)
			mock.ExpectQuery(query).
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "user_id", "classroom_id", "document_id"}).
					AddRow(notes[0].Id, notes[0].UserId, notes[0].ClassroomId, notes[0].DocumentId).
					AddRow(notes[1].Id, notes[1].UserId, notes[1].ClassroomId, notes[1].DocumentId))

			listNotesV1Response, err = grpcApi.ListNotesV1(ctx, listNotesV1Request)
		})

		It("successful retrieval of the list of notes", func() {
			Expect(err).Should(BeNil())
			Expect(listNotesV1Response.Notes[0].Id).Should(Equal(int64(notes[0].Id)))
			Expect(listNotesV1Response.Notes[0].UserId).Should(Equal(int32(notes[0].UserId)))
			Expect(listNotesV1Response.Notes[0].ClassroomId).Should(Equal(int32(notes[0].ClassroomId)))
			Expect(listNotesV1Response.Notes[0].DocumentId).Should(Equal(int32(notes[0].DocumentId)))

			Expect(listNotesV1Response.Notes[1].Id).Should(Equal(int64(notes[1].Id)))
			Expect(listNotesV1Response.Notes[1].UserId).Should(Equal(int32(notes[1].UserId)))
			Expect(listNotesV1Response.Notes[1].ClassroomId).Should(Equal(int32(notes[1].ClassroomId)))
			Expect(listNotesV1Response.Notes[1].DocumentId).Should(Equal(int32(notes[1].DocumentId)))
		})
	})

	Context("remove note with invalid arguments", func() {

		BeforeEach(func() {
			removeNoteV1Request = &desc.RemoveNoteV1Request{
				NoteId: -1,
			}

			// setting the wait for the mock request is not required,
			// since the error will return earlier due to invalid arguments

			removeNoteV1Response, err = grpcApi.RemoveNoteV1(ctx, removeNoteV1Request)
		})

		It("failed note removal due to invalid arguments", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(removeNoteV1Response).Should(BeNil())
		})
	})

	Context("unsuccessful note removal", func() {

		BeforeEach(func() {
			removeNoteV1Request = &desc.RemoveNoteV1Request{
				NoteId: 1,
			}

			mock.ExpectExec("DELETE FROM notes").
				WithArgs(removeNoteV1Request.NoteId).
				WillReturnError(errors.New("failed to execute sql request"))

			removeNoteV1Response, err = grpcApi.RemoveNoteV1(ctx, removeNoteV1Request)
		})

		It("failed to execute sql request", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(removeNoteV1Response).Should(BeNil())
		})
	})

	Context("remove note", func() {

		BeforeEach(func() {
			removeNoteV1Request = &desc.RemoveNoteV1Request{
				NoteId: 1,
			}

			mock.ExpectExec("DELETE FROM notes").
				WithArgs(removeNoteV1Request.NoteId).
				WillReturnResult(sqlmock.NewResult(0, 1))

			dataProducerMock.EXPECT().Send(gomock.Any())

			removeNoteV1Response, err = grpcApi.RemoveNoteV1(ctx, removeNoteV1Request)
		})

		It("successful removal of a note in the database", func() {
			Expect(err).Should(BeNil())
			Expect(removeNoteV1Response.Found).Should(Equal(true))
		})
	})
})
