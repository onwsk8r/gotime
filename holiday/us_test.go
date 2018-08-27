package holiday_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/gotime/holiday"
)

var _ = Describe("US Holidays", func() {
	year := 2018

	// It("parseYear should default to the current year", func() {
	// 	Expect(parseYear()).To(Equal(time.Now().Year()))
	// })

	It("Should find New Year's Eve", func() {
		Expect(NYEve(year).Format("20060102")).To(Equal("20171231"))
	})

	It("Should find New Year's Day", func() {
		Expect(NYDay(year).Format("20060102")).To(Equal("20180101"))
	})

	// Inauguration day
	It("should find inauguration day", func() {
		Expect(InaugurationDay(year).Year()).To(Equal(2021))
		Expect(InaugurationDay().YearDay()).To(Equal(20))
	})

	It("Should find MLK Day", func() {
		Expect(MLKDay(year).Format("20060102")).To(Equal("20180115"))
	})

	It("Should find President's Day", func() {
		Expect(PresidentsDay(year).Format("20060102")).To(Equal("20180219"))
	})

	It("Should find Memorial Day", func() {
		Expect(MemorialDay(year).Format("20060102")).To(Equal("20180528"))
	})

	It("Should find Independence Day", func() {
		Expect(IndependenceDay(year).Format("20060102")).To(Equal("20180704"))
	})

	It("Should find Labor Day", func() {
		Expect(LaborDay(year).Format("20060102")).To(Equal("20180903"))
	})

	It("Should find Columbus Day", func() {
		Expect(ColumbusDay(year).Format("20060102")).To(Equal("20181008"))
	})

	It("Should find Veteran's Day", func() {
		Expect(VeteransDay(year).Format("20060102")).To(Equal("20181111"))
	})

	It("Should find Thanksgiving Day", func() {
		Expect(Thanksgiving(year).Format("20060102")).To(Equal("20181122"))
	})

	It("Should find Black  Friday", func() {
		Expect(BlackFriday(year).Format("20060102")).To(Equal("20181123"))
	})

	It("Should find Christmas Eve", func() {
		Expect(ChristmasEve(year).Format("20060102")).To(Equal("20181224"))
	})

	It("Should find Christmas Day", func() {
		Expect(ChristmasDay(year).Format("20060102")).To(Equal("20181225"))
	})
})
