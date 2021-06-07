package alarmer_test

import (
	"sync/atomic"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-note-api/core/alarmer"
)

var _ = Describe("Alarmer", func() {

	var (
		alrm  alarmer.Alarmer
		timer *time.Timer
	)

	BeforeEach(func() {
		alrm = alarmer.New(20 * time.Millisecond)
		timer = time.NewTimer(500 * time.Millisecond)

		err := alrm.Init()

		if err != nil {
			Fail("alarmer initialization failed")
		}
	})

	Context("Alarm frequency", func() {
		It("closing alarm", func() {

			timer.Reset(0)

			go func() {
				defer alrm.Close()
				<-timer.C
			}()

			Eventually(alrm.Alarm()).Should(BeClosed())
		})

		It("number of alarms", func() {

			timer.Reset(0)

			var count uint32
			timer := time.NewTimer(2 * time.Second)

			go func() {
				defer alrm.Close()

				<-timer.C

				//fmt.Printf("count: %v\n", atomic.LoadUint32(&count))
				Expect(atomic.LoadUint32(&count)).To(SatisfyAll(BeNumerically(">=", 90), BeNumerically("<=", 105)))
			}()

			for range alrm.Alarm() {
				atomic.AddUint32(&count, 1)
			}
		})
	})
})
