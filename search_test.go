package gotime_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/gotime"
)

// nolint: dupl
var _ = Describe("Search Functions", func() {
	Describe("NthWeekday", func() {
		It("should find MLK Day", func() {
			exp := time.Date(2018, time.January, 15, 0, 0, 0, 0, time.UTC)
			res := NthWeekday(2018, time.January, 3, time.Monday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180115"))
		})
		It("should find the day this test was written", func() {
			exp := time.Date(2018, time.August, 26, 0, 0, 0, 0, time.UTC)
			res := NthWeekday(2018, time.August, 4, time.Sunday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180826"))
		})
		It("should find Thanksgiving", func() {
			exp := time.Date(2018, time.November, 22, 0, 0, 0, 0, time.UTC)
			res := NthWeekday(2018, time.November, 4, time.Thursday)
			fmt.Println(res.String())
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20181122"))
		})
	})

	Describe("FirstWeekday", func() {
		It("should find the first Saturday of a month that starts on Saturday", func() {
			exp := time.Date(2018, time.September, 1, 0, 0, 0, 0, time.UTC)
			res := FirstWeekday(2018, time.September, time.Saturday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180901"))
		})
		It("should find the first Friday of a month that starts on Saturday", func() {
			exp := time.Date(2018, time.September, 7, 0, 0, 0, 0, time.UTC)
			res := FirstWeekday(2018, time.September, time.Friday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180907"))
		})
	})

	Describe("LastWeekday", func() {
		It("should find the last Sunday of a month that starts on Saturday", func() {
			exp := time.Date(2018, time.September, 30, 0, 0, 0, 0, time.UTC)
			res := LastWeekday(2018, time.September, time.Sunday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180930"))
		})
		It("should find the last Friday of a month that starts on Saturday", func() {
			exp := time.Date(2018, time.September, 28, 0, 0, 0, 0, time.UTC)
			res := LastWeekday(2018, time.September, time.Friday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180928"))
		})
		It("should find the last Monday of a month that starts on Saturday", func() {
			exp := time.Date(2018, time.September, 24, 0, 0, 0, 0, time.UTC)
			res := LastWeekday(2018, time.September, time.Monday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180924"))
		})
		It("should find the last day of a 31-day month", func() {
			exp := time.Date(2018, time.August, 31, 0, 0, 0, 0, time.UTC)
			res := LastWeekday(2018, time.August, time.Friday)
			Expect(res.Equal(exp)).To(BeTrue())
			Expect(res.Format("20060102")).To(Equal("20180831"))
		})
	})
})

func ExampleNthWeekday() {
	year := 2018
	month := time.September
	n := 3
	day := time.Monday

	res := NthWeekday(year, month, n, day)
	fmt.Println("The third Monday in September is", res.Format("2006-01-02"))
	// Output: The third Monday in September is 2018-09-17
}
