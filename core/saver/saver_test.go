package saver_test

import (
	"context"
	"sync"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-note-api/core/alarmer"
	"github.com/ozoncp/ocp-note-api/core/flusher"
	"github.com/ozoncp/ocp-note-api/core/mocks"
	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/saver"
)

var _ = Describe("Saver", func() {

	var (
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockRepo
		flusherTest flusher.Flusher
		alarmerTest alarmer.Alarmer
		saverTest   saver.Saver
		capacity    uint32
		chunkSize   uint32
		duration    time.Duration

		ctx context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		duration = 100 * time.Millisecond
		capacity = 8
		chunkSize = 5

		flusherTest = flusher.New(mockRepo, chunkSize)
		alarmerTest = alarmer.New(duration)

		ctx = context.Background()
		saverTest = saver.New(ctx, capacity, flusherTest, alarmerTest, true)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("creating a new saver", func() {

		When("invalid input data", func() {
			It("capacity = 0", func() {
				alarmerTest = alarmer.New(duration)
				saverTest = saver.New(ctx, 0, flusherTest, alarmerTest, true)

				Expect(saverTest).Should(BeNil())
			})

			It("flusher = nil", func() {
				alarmerTest = alarmer.New(duration)
				saverTest = saver.New(ctx, capacity, nil, alarmerTest, true)

				Expect(saverTest).Should(BeNil())
			})

			It("alarmer = nil", func() {
				alarmerTest = alarmer.New(0)
				saverTest = saver.New(ctx, capacity, flusherTest, alarmerTest, true)

				Expect(saverTest).Should(BeNil())
			})

			It("all input parameters are invalid", func() {
				alarmerTest = alarmer.New(0)
				saverTest = saver.New(ctx, 0, nil, alarmerTest, true)

				Expect(saverTest).Should(BeNil())
			})
		})

		When("correct input", func() {
			It("all input parameters are correct", func() {
				alarmerTest = alarmer.New(duration)
				saverTest = saver.New(ctx, capacity, flusherTest, alarmerTest, true)

				Expect(saverTest).ShouldNot(BeNil())
			})
		})
	})

	Context("saver initialization", func() {

		When("correct input", func() {
			It("correct initialization", func() {
				err := saverTest.Init()
				Expect(err).Should(BeNil())
				saverTest.Close()
			})
		})

		When("the saver has not been initialized", func() {
			It("duplicate initialization", func() {
				err := saverTest.Init()
				Expect(err).Should(BeNil())

				err = saverTest.Init()
				Expect(err).ShouldNot(BeNil())
				saverTest.Close()
			})
		})
	})

	Context("goroutine work", func() {

		var (
			notesNum uint32
			chunkNum int
		)

		When("correct input", func() {

			It("the number of notes is less than the storage size", func() {
				err := saverTest.Init()

				if err != nil {
					Fail("saver initialization failed")
				}

				notesNum = capacity - 3
				chunkNum = int(notesNum / chunkSize)

				if notesNum%chunkSize != 0 {
					chunkNum++
				}

				var wg sync.WaitGroup
				wg.Add(chunkNum)

				mockRepo.EXPECT().MultiAddNotes(ctx, gomock.Any()).AnyTimes().Do(func(ctx context.Context, notes []note.Note) {
					wg.Done()
				}).Return(uint64(0), nil)

				for i := 0; i < int(notesNum); i++ {
					saverTest.Save(note.Note{
						Id:          uint64(i + 1),
						UserId:      0,
						ClassroomId: 0,
						DocumentId:  0,
					})
				}

				wg.Wait()
				saverTest.Close()
			})

			It("the number of notes is larger than the storage size (\"delete all\" mode)", func() {

				err := saverTest.Init()

				if err != nil {
					Fail("saver initialization failed")
				}

				notesNum = capacity + 3
				chunkNum = int((notesNum - capacity) / chunkSize)

				if notesNum%chunkSize != 0 {
					chunkNum++
				}

				var wg sync.WaitGroup
				wg.Add(chunkNum)

				mockRepo.EXPECT().MultiAddNotes(ctx, gomock.Any()).AnyTimes().Do(func(ctx context.Context, notes []note.Note) {
					wg.Done()
				}).Return(uint64(0), nil)

				for i := 0; i < int(notesNum); i++ {
					saverTest.Save(note.Note{
						Id:          uint64(i * 2),
						UserId:      0,
						ClassroomId: 0,
						DocumentId:  0,
					})
				}

				wg.Wait()
				saverTest.Close()
			})

			It("the number of notes is larger than the storage size (\"delete first\" mode)", func() {

				saverTest := saver.New(ctx, capacity, flusherTest, alarmerTest, false)

				err := saverTest.Init()

				if err != nil {
					Fail("saver initialization failed")
				}

				notesNum = capacity + 4
				chunkNum = int(capacity / chunkSize)

				if notesNum%chunkSize != 0 {
					chunkNum++
				}

				var wg sync.WaitGroup
				wg.Add(chunkNum)

				mockRepo.EXPECT().MultiAddNotes(ctx, gomock.Any()).AnyTimes().Do(func(ctx context.Context, notes []note.Note) {
					wg.Done()
				}).Return(uint64(0), nil)

				for i := 0; i < int(notesNum); i++ {
					saverTest.Save(note.Note{
						Id:          uint64(i * 3),
						UserId:      0,
						ClassroomId: 0,
						DocumentId:  0,
					})
				}

				wg.Wait()
				saverTest.Close()
			})
		})

		When("saver initialization not performed", func() {
			It("return panic when creating saver", func() {

				Expect(func() {
					saverTest.Save(note.Note{})
				}).Should(Panic())
			})
		})

		When("flush data when closing saver", func() {
			It("flush data", func() {

				err := saverTest.Init()

				if err != nil {
					Fail("saver initialization failed")
				}
				notesNum = capacity - 3
				chunkNum = int(notesNum / chunkSize)

				if notesNum%chunkSize != 0 {
					chunkNum++
				}

				var wg sync.WaitGroup
				wg.Add(chunkNum)

				mockRepo.EXPECT().MultiAddNotes(ctx, gomock.Any()).AnyTimes().Do(func(ctx context.Context, notes []note.Note) {
					wg.Done()
				}).Return(uint64(0), nil)

				for i := 0; i < int(notesNum); i++ {
					saverTest.Save(note.Note{
						Id:          uint64(i * 2),
						UserId:      0,
						ClassroomId: 0,
						DocumentId:  0,
					})
				}

				saverTest.Close()
				wg.Wait()
			})
		})
	})
})
