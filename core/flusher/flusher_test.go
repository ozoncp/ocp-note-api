package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-note-api/core/flusher"
	"github.com/ozoncp/ocp-note-api/core/mocks"
	"github.com/ozoncp/ocp-note-api/core/note"
)

var (
	errDeadlineExceeded = errors.New("mock error")
)

var _ = Describe("Flusher", func() {
	var (
		err error

		ctrl *gomock.Controller

		mockStorage *mocks.MockRepo

		notes  []note.Note
		result []note.Note

		f flusher.Flusher

		chunkSize int
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockStorage = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.New(mockStorage, chunkSize)
		result = f.Flush(notes)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo saves all notes", func() {
		BeforeEach(func() {
			notes = []note.Note{{}}
			chunkSize = 2

			mockStorage.EXPECT().AddNotes(gomock.Any()).Return(nil).MinTimes(1)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(result).Should(BeNil())
		})
	})

	Context("repo don't saves any note", func() {
		BeforeEach(func() {
			notes = []note.Note{{}, {}}
			chunkSize = 2

			mockStorage.EXPECT().AddNotes(gomock.Any()).Return(errDeadlineExceeded)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(result).Should(BeEquivalentTo(notes))
		})
	})

	Context("repo saves half notes", func() {
		BeforeEach(func() {
			notes = []note.Note{{}, {}}
			chunkSize = len(notes) / 2

			gomock.InOrder(
				mockStorage.EXPECT().AddNotes(gomock.Any()).Return(nil),
				mockStorage.EXPECT().AddNotes(gomock.Any()).Return(errDeadlineExceeded),
			)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(result).Should(BeEquivalentTo(notes[chunkSize:]))
		})
	})
})
