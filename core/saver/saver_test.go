package saver_test

import (
	"errors"
	"sync"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-note-api/core/alarmer"
	"github.com/ozoncp/ocp-note-api/core/flusher"
	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/repo"
	"github.com/ozoncp/ocp-note-api/core/saver"
)

var _ = Describe("Saver", func() {

	var (
		ctrl    *gomock.Controller
		storage repo.Repo
		flush   flusher.Flusher
		alarm   alarmer.Alarmer
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		storage = &repoStub{}
		flush = flusher.New(storage, 5)
		alarm = alarmer.New(5 * time.Second)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Init", func() {
		It("periodic save", func() {
			notes := []note.Note{
				{Id: 1},
				{Id: 2},
				{Id: 3},
				{Id: 4},
				{Id: 5},
			}

			var wg sync.WaitGroup
			wg.Add(1)
			defer wg.Wait()

			saver := saver.New(uint(len(notes)+1), flush, alarm, true)

			go func() {
				defer saver.Close()
				defer wg.Done()

				saver.Init()

				for i := 0; i < len(notes); i++ {
					saver.Save(notes[i])
				}
			}()
		})
	})
})

type repoStub struct {
	notes []note.Note
}

func (r *repoStub) AddNotes(notes []note.Note) error {
	r.notes = append(r.notes, notes...)
	return nil
}

func (r *repoStub) DescribeNote(id uint64) (*note.Note, error) {
	for _, val := range r.notes {
		if val.Id == id {
			return &val, nil
		}
	}

	return nil, errors.New("id not found")
}

func (r *repoStub) ListNotes(count, offset uint64) ([]note.Note, error) {
	return r.notes[offset : offset+count], nil
}

func (r *repoStub) RemoveNote(id uint64) error {
	return nil
}
