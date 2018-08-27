package holiday_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHoliday(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Holiday Suite")
}
