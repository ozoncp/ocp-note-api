package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-note-api/core/flusher"
	"github.com/ozoncp/ocp-note-api/core/mocks"
	"github.com/ozoncp/ocp-note-api/core/note"
)

var _ = Describe("Flusher", func() {

	var (
		ctrl *gomock.Controller

		mockStorage *mocks.MockRepo

		//notes []note.Note
		notes []note.Note
		rest  []note.Note

		f         flusher.Flusher
		chunkSize int
		err       error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockStorage = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.New(chunkSize, mockStorage)
		rest = f.Flush(notes)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all notes", func() {
		BeforeEach(func() {
			//notes = []note.Note{{}}

			mockStorage.EXPECT().AddNotes(gomock.Any()).Return(nil).MinTimes(1)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})

		AfterEach(func() {

		})
	})
})
