package gotime_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onwsk8r/gotime"
)

func init() {
	// Make sure "local" time is CDT
	var err error
	if time.Local, err = time.LoadLocation("America/Chicago"); err != nil {
		Fail("Could not set time.Local")
	}
}

// nolint: dupl
var _ = Describe("Equality Functions", func() {
	var refLocal, refUTC, refDateLocal, refTimeLocal, refDateUTC, refTimeUTC time.Time

	BeforeEach(func() {
		// Give us some times to play with
		refLocal = time.Date(2006, time.January, 2, 15, 4, 5, 0, time.Local)
		refUTC = time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)

		refDateLocal = time.Date(2006, time.January, 2, 3, 56, 24, 55, time.Local)
		refTimeLocal = time.Date(2013, time.October, 25, 15, 4, 5, 0, time.Local)
		refDateUTC = time.Date(2006, time.January, 2, 3, 56, 24, 55, time.UTC)
		refTimeUTC = time.Date(2013, time.October, 25, 15, 4, 5, 0, time.UTC)
	})

	Describe("DateEquals function", func() {
		It("should return true with identical values", func() {
			Expect(DateEquals(refLocal, refLocal)).To(BeTrue())
		})
		It("should return true if only the time zones are different", func() {
			Expect(DateEquals(refLocal, refUTC)).To(BeTrue())
		})
		It("should return true if the times are different", func() {
			Expect(DateEquals(refLocal, refDateLocal)).To(BeTrue())
		})
		It("should return true if the times and TZs are different", func() {
			Expect(DateEquals(refDateUTC, refDateLocal)).To(BeTrue())
		})
		It("should return false if the dates are different", func() {
			Expect(DateEquals(refTimeUTC, refUTC)).To(BeFalse())
		})
		It("should return false if the dates are different because of the TZs", func() {
			Expect(DateEquals(refDateUTC, refDateUTC.In(time.Local))).To(BeFalse())
		})
	})

	Describe("DateEqualsTZ function", func() {
		It("should return true with identical values", func() {
			Expect(DateEqualsTZ(refLocal, refLocal)).To(BeTrue())
		})
		It("should return true if only the time zones are different", func() {
			Expect(DateEqualsTZ(refLocal, refLocal.In(time.UTC))).To(BeTrue())
		})
		It("should return true if the times are different", func() {
			Expect(DateEqualsTZ(refLocal, refDateLocal)).To(BeTrue())
		})
		It("should return true if the times and TZs are different", func() {
			Expect(DateEqualsTZ(refDateUTC, refDateUTC)).To(BeTrue())
		})
		It("should return false if the dates are different", func() {
			Expect(DateEqualsTZ(refTimeUTC, refUTC)).To(BeFalse())
		})
	})

	Describe("TimeEquals function", func() {
		It("should return true with identical values", func() {
			Expect(TimeEquals(refLocal, refLocal)).To(BeTrue())
		})
		It("should return true if only the time zones are different", func() {
			Expect(TimeEquals(refLocal, refUTC)).To(BeTrue())
		})
		It("should return true if the dates are different", func() {
			Expect(TimeEquals(refLocal, refTimeLocal)).To(BeTrue())
		})
		It("should return true if the dates and TZs are different", func() {
			Expect(TimeEquals(refTimeUTC, refLocal)).To(BeTrue())
		})
		It("should return false if the times are different", func() {
			Expect(TimeEquals(refDateUTC, refUTC)).To(BeFalse())
		})
		It("should return false if the times are different because of the TZs", func() {
			Expect(TimeEquals(refLocal, refLocal.In(time.UTC))).To(BeFalse())
		})
	})

	Describe("TimeEqualsTZ function", func() {
		It("should return true with identical values", func() {
			Expect(TimeEqualsTZ(refLocal, refLocal)).To(BeTrue())
		})
		It("should return true if the times are only different because of the TZs", func() {
			Expect(TimeEqualsTZ(refLocal, refLocal.In(time.UTC))).To(BeTrue())
		})
		It("should return true if the dates are different", func() {
			// Add one hour because of DST
			Expect(TimeEqualsTZ(refLocal, refTimeLocal.Add(time.Hour))).To(BeTrue())
		})
		It("should return true if the dates are different but in UTC", func() {
			Expect(TimeEqualsTZ(refUTC, refTimeUTC)).To(BeTrue())
		})
		It("should return false if the times are different", func() {
			Expect(TimeEqualsTZ(refDateUTC, refUTC)).To(BeFalse())
		})
		It("should return false if only the time zones are different", func() {
			Expect(TimeEqualsTZ(refLocal, refUTC)).To(BeFalse())
		})
	})

	// These next three functions don't need much testing because they just
	// call other functions; they are only for convenience.
	Describe("TimeEqualsNS function", func() {
		It("should return true with identical values", func() {
			Expect(TimeEqualsTZ(refLocal, refLocal)).To(BeTrue())
		})
	})

	Describe("TimeEqualsNSTZ function", func() {
		It("should return true with identical values", func() {
			Expect(TimeEqualsNSTZ(refLocal, refLocal)).To(BeTrue())
		})
		It("should return false with different values", func() {
			Expect(SameTime(refLocal, refUTC)).To(BeFalse())
		})
	})

	Describe("SameTime function", func() {
		It("should return true with identical values", func() {
			Expect(SameTime(refLocal, refLocal)).To(BeTrue())
		})
		It("should return false with different values", func() {
			Expect(SameTime(refLocal, refUTC)).To(BeFalse())
		})
	})
})
