package gotime_test

import (
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/gotime"
)

var timestamps = []string{"2006-01-02T15:04:05", "20060102T150405",
	"2006-01-02T15:04:05Z", "20060102T150405-07:00", "2006-01-02T15:04:05.000"}

var _ = Describe("Parse", func() {
	Describe("Parse()", func() {
		It("should parse a timestamp correctly", func() {
			now := time.Now()
			Y, M, D := now.Date()
			h, m, s := now.Clock()
			str := fmt.Sprintf("%04d%02d%02dT%02d%02d%02d", Y, M, D, h, m, s)
			res, err := Parse(str)
			Expect(err).ToNot(HaveOccurred())
			Expect(now.Format("20060102030405")).To(Equal(res.Format("20060102030405")))
		})

		It("should return an error if DateParser does", func() {
			// Note that GetTimeFormatFast does not return any errors
			res, err := Parse("----T03:04:05")
			Expect(res).To(BeZero())
			Expect(err).To(HaveOccurred())
		})
	})

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

		// Exceptions to the rule
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
		// Pretty sure that "Z07:00" is far from valid, but according
		// to Google it is, like, some sort of standard.
		var times = []string{"T15:04:05", "T150405", "T1504", "T15:04",
			"T15", "T150405Z", "T150405-07:00", "T150405Z07:00"}
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
