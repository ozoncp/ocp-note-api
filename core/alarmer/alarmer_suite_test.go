package alarmer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAlarmer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alarmer Suite")
}
