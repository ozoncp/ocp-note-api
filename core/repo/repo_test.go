package repo_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-note-api/core/mocks"
	"github.com/ozoncp/ocp-note-api/core/note"
)

var _ = Describe("Repo", func() {

	var (
		ctrl *gomock.Controller

		mockRepo *mocks.MockRepo

		notes []note.Note

		//r repo.Repo
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		// r =
		// rest = r.AddNotes()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all notes", func() {
		BeforeEach(func() {
			notes = []note.Note{{}}

			mockRepo.EXPECT().AddNotes(gomock.Any()).Return(nil).MinTimes(1)
		})

		AfterEach(func() {

		})
	})
})
