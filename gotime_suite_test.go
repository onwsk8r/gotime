package gotime_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGotime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gotime Suite")
}
