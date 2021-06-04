package alarmer_test

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-note-api/core/alarmer"
)

var _ = Describe("Alarmer", func() {

	var (
		ctrl *gomock.Controller
		alrm alarmer.Alarmer
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		alrm = alarmer.New(5 * time.Millisecond)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Alarm frequency", func() {
		It("closing alarm", func() {
			alrm.Init()

			timer := time.NewTimer(250 * time.Millisecond)

			go func() {
				defer alrm.Close()
				<-timer.C
			}()

			Eventually(alrm.Alarm()).Should(BeClosed())
		})
	})
})
