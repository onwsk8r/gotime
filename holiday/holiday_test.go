package holiday_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/gotime/holiday"
)

var _ = Describe("Holiday", func() {
	It("should return observed holidays", func() {
		easter := Easter()
		Expect(Observed(easter)).To(Equal(easter.Add(24 * time.Hour)))
	})

	Describe("List type", func() {
		It("should determine whether a date is an exact holiday", func() {
			theDay, _ := time.Parse("20060102", "20180101")
			Expect(FederalHolidays.Contains(theDay)).To(BeTrue())
		})

		It("should determine whether a date is an observed holiday", func() {
			theDay, _ := time.Parse("20060102", "20180101")
			Expect(FederalHolidays.Observes(theDay)).To(BeTrue())
		})
	})
})
