package gotime_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/gotime"
)

var timestamps = []string{"2006-01-02T15:04:05", "20060102T150405",
	"2006-01-02T15:04:05Z", "20060102T150405-07:00", "2006-01-02T15:04:05.000"}

var _ = Describe("Parse", func() {
	Describe("GetDateFormatFast", func() {
		var dates = []string{"2006-01-02", "2006-01", "20060102", "--0102", "--01-02"}
		for _, d := range dates {
			It(fmt.Sprintf("It should parse date %s", d), func() {
				res, err := GetDateFormatFast(d)
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(d))
			})
		}
		for _, d := range timestamps { // nolint: dupl
			It(fmt.Sprintf("It should parse timestamp %s", d), func() {
				res, err := GetDateFormatFast(d)
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(strings.Split(d, "T")[0]))
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

	Describe("GetTimeFormatFast", func() {
		var times = []string{"15:04:05", "150405", "1504", "15:04", "15", "150405Z", "150405-07:00"}
		for _, d := range times {
			It(fmt.Sprintf("It should parse time %s", d), func() {
				res, err := GetTimeFormatFast(d)
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(d))
			})
		}

		for _, d := range timestamps { // nolint: dupl
			It(fmt.Sprintf("It should parse timestamp %s", d), func() {
				res, err := GetDateFormatFast(d)
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(strings.Split(d, "T")[0]))
			})
		}
	})
})
