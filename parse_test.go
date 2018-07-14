package gotime_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/gotime"
)

var _ = Describe("Parse", func() {
	var dates = []string{"2006-01-02", "2006-01", "20060102", "--0102", "--01-02"}
	Describe("GetDateFormatFast", func() {
		for _, d := range dates {
			It(fmt.Sprintf("It should parse %s", d), func() {
				res, err := GetDateFormatFast(d)
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(d))
			})
		}
		It("should choke on week numbers", func() {
			res, err := GetDateFormatFast("2006-W01")
			Expect(res).To(BeZero())
			Expect(err).To(BeAssignableToTypeOf(&ParseError{}))
			Expect(err.Error()).To(ContainSubstring("Cannot parse week numbers"))
		})
		It("should choke on day numbers with hyphens", func() {
			res, err := GetDateFormatFast("2006-002")
			Expect(res).To(BeZero())
			Expect(err).To(BeAssignableToTypeOf(&ParseError{}))
			Expect(err.Error()).To(ContainSubstring("Cannot parse day of year"))
		})
		It("should choke on day numbers without hyphens", func() {
			res, err := GetDateFormatFast("2006002")
			Expect(res).To(BeZero())
			Expect(err).To(BeAssignableToTypeOf(&ParseError{}))
			Expect(err.Error()).To(ContainSubstring("Cannot parse day of year"))
		})
	})
})
