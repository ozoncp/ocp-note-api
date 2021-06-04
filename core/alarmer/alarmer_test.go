package alarmer_test

import (
	"fmt"
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
	})

	Context("Alarm frequency", func() {
		It("closing alarm", func() {
			alrm.Init()
			timer.Reset(0)

			go func() {
				defer alrm.Close()
				<-timer.C
			}()

			Eventually(alrm.Alarm()).Should(BeClosed())
		})

		It("number of alarms", func() {
			alrm.Init()
			timer.Reset(0)

			var count uint32
			timer := time.NewTimer(2000 * time.Millisecond)

			go func() {
				defer alrm.Close()

				<-timer.C

				fmt.Printf("count: %v\n", atomic.LoadUint32(&count))
				Expect(atomic.LoadUint32(&count)).To(SatisfyAll(BeNumerically(">=", 99), BeNumerically("<=", 101)))
			}()

			for range alrm.Alarm() {
				atomic.AddUint32(&count, 1)
			}
		})
	})
})
