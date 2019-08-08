package main

import (
	"time"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type picoyplacaSuite struct{}

var _ = Suite(&picoyplacaSuite{})

// Unit tests

func (s *picoyplacaSuite) Test_parseDateTime_RFC3339_ValidDateValidTime(c *C) {
	validDateString := "2019-05-01"
	validTimeString := "07:35:41"

	expectedWeekDay := time.Wednesday
	expectedDate := validDateString
	expectedTime := validTimeString

	actualDateTime, actualError := parseDateTime(validDateString, validTimeString)
	actualDate := actualDateTime.Format("2006-01-02")
	actualTime := actualDateTime.Format("15:04:05")
	actualWeekDay := actualDateTime.Weekday()

	c.Assert(actualError, IsNil)
	c.Assert(actualDate, Equals, expectedDate)
	c.Assert(actualTime, Equals, expectedTime)
	c.Assert(actualWeekDay, Equals, expectedWeekDay)
}

func (s *picoyplacaSuite) Test_parseDateTime_RFC3339_ValidDateInvalidTime(c *C) {
	validDateString := "2019-05-01"
	invalidTimeString := "35:80:90"

	expectedWeekDay := time.Monday
	expectedDate := "0001-01-01"
	expectedTime := "00:00:00"

	actualDateTime, actualError := parseDateTime(validDateString, invalidTimeString)
	actualDate := actualDateTime.Format("2006-01-02")
	actualTime := actualDateTime.Format("15:04:05")
	actualWeekDay := actualDateTime.Weekday()

	c.Assert(actualError, NotNil)
	c.Assert(actualError, ErrorMatches, "Invalid Date or Time given.")
	c.Assert(actualDate, Equals, expectedDate)
	c.Assert(actualTime, Equals, expectedTime)
	c.Assert(actualWeekDay, Equals, expectedWeekDay)
}

func (s *picoyplacaSuite) Test_parseDateTime_RFC3339_InvalidDateValidTime(c *C) {
	invalidDateString := "20-05-2019"
	validTimeString := "15:00:00"

	expectedWeekDay := time.Monday
	expectedDate := "0001-01-01"
	expectedTime := "00:00:00"

	actualDateTime, actualError := parseDateTime(invalidDateString, validTimeString)
	actualDate := actualDateTime.Format("2006-01-02")
	actualTime := actualDateTime.Format("15:04:05")
	actualWeekDay := actualDateTime.Weekday()

	c.Assert(actualError, NotNil)
	c.Assert(actualError, ErrorMatches, "Invalid Date or Time given.")
	c.Assert(actualDate, Equals, expectedDate)
	c.Assert(actualTime, Equals, expectedTime)
	c.Assert(actualWeekDay, Equals, expectedWeekDay)
}

func (s *picoyplacaSuite) Test_parseDateTime_RFC3339_InvalidDateInvalidTime(c *C) {
	invalidDateString := "2019-15-45"
	invalidTimeString := "30:80:90"

	expectedWeekDay := time.Monday
	expectedDate := "0001-01-01"
	expectedTime := "00:00:00"

	actualDateTime, actualError := parseDateTime(invalidDateString, invalidTimeString)
	actualDate := actualDateTime.Format("2006-01-02")
	actualTime := actualDateTime.Format("15:04:05")
	actualWeekDay := actualDateTime.Weekday()

	c.Assert(actualError, NotNil)
	c.Assert(actualError, ErrorMatches, "Invalid Date or Time given.")
	c.Assert(actualDate, Equals, expectedDate)
	c.Assert(actualTime, Equals, expectedTime)
	c.Assert(actualWeekDay, Equals, expectedWeekDay)
}

func (s *picoyplacaSuite) Test_validateLicensePlate_ValidLicensePate(c *C) {
	validLicensePlate := "AAC-1111"

	expectedValidity := true

	actualValidity := validateLicensePlate(validLicensePlate)

	c.Assert(actualValidity, Equals, expectedValidity)
}

func (s *picoyplacaSuite) Test_validateLicensePlate_MissingHyphen(c *C) {
	validLicensePlate := "AAC1111"

	expectedValidity := false

	actualValidity := validateLicensePlate(validLicensePlate)

	c.Assert(actualValidity, Equals, expectedValidity)
}

func (s *picoyplacaSuite) Test_validateLicensePlate_NonAlphabetBeforeHyphen(c *C) {
	validLicensePlate := "AAC1111"

	expectedValidity := false

	actualValidity := validateLicensePlate(validLicensePlate)

	c.Assert(actualValidity, Equals, expectedValidity)
}

func (s *picoyplacaSuite) Test_validateLicensePlate_NonNumericAfterHyphen(c *C) {
	validLicensePlate := "AAC-1A11"

	expectedValidity := false

	actualValidity := validateLicensePlate(validLicensePlate)

	c.Assert(actualValidity, Equals, expectedValidity)
}

func (s *picoyplacaSuite) Test_extractLastDigit_ValidLicensePlate(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_extractLastDigit_InvalidLicensePlate(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_1(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_2(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_3(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_4(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_5(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitForRespectiveDays_1(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitForRespectiveDays_2(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_3(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_4(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_5(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_6(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_7(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_8(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_9(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_Weekend_1(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}

func (s *picoyplacaSuite) Test_allowedInCity_Weekend_2(c *C) {
	c.Skip("XXX: Not implemented")
	c.Fail()
}
