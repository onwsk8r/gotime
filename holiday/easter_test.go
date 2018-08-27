package holiday_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/gotime/holiday"
)

var _ = Describe("Easter", func() {
	It("Should find Easter", func() {
		Expect(Easter(2019).Format("20060102")).To(Equal("20190421"))
	})

	It("Should find Good Friday", func() {
		Expect(GoodFriday(2018).Format("20060102")).To(Equal("20180330"))
	})
})
