package main

import (
	"os"
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

func (s *picoyplacaSuite) Test_validateLicensePlate_ValidLicensePlate(c *C) {
	validLicensePlate := "AAC-1111"

	expectedValidity := true

	actualValidity := validateLicensePlate(validLicensePlate)

	c.Assert(actualValidity, Equals, expectedValidity)
}

func (s *picoyplacaSuite) Test_validateLicensePlate_ValidLicensePlateOld(c *C) {
	validLicensePlate := "AAC-111"

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

func (s *picoyplacaSuite) Test_validateLicensePlate_InvalidAlphabetLength(c *C) {
	validLicensePlate := "ABCD-1234"

	expectedValidity := false

	actualValidity := validateLicensePlate(validLicensePlate)

	c.Assert(actualValidity, Equals, expectedValidity)
}

func (s *picoyplacaSuite) Test_validateLicensePlate_InvalidDigitLength(c *C) {
	validLicensePlate := "ABC-12345"

	expectedValidity := false

	actualValidity := validateLicensePlate(validLicensePlate)

	c.Assert(actualValidity, Equals, expectedValidity)
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_1(c *C) {
	lastDigit := 1
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-06T06:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

// These set of Unit tests, tests if the given last digit works as expected
// for different times of day restricted / unrestricted times etc
func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_2(c *C) {
	lastDigit := 1
	// Restricted Morning
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-06T07:30:00-05:00")

	expectedResult := false

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_3(c *C) {
	lastDigit := 1
	// Unrestricted Morning
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-06T10:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_4(c *C) {
	lastDigit := 1
	// Restricted Evening
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-06T17:30:00-05:00")

	expectedResult := false

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_FixedDigitDifferentTimes_5(c *C) {
	lastDigit := 1
	// Unrestricted Evening
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-06T20:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

// These set of unit tests, tests if the conditional check for the various last
// digits for the respective days of week work as expected
func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitForRespectiveDays_1(c *C) {
	// Monday + Restricted Evening
	lastDigit := 2
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-06T16:30:00-05:00")

	expectedResult := false

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitForRespectiveDays_2(c *C) {
	// Tuesday + Restricted Morning
	lastDigit := 3
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-07T07:30:00-05:00")

	expectedResult := false

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_3(c *C) {
	// Tuesday + Unrestricted Morning
	lastDigit := 4
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-07T10:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_4(c *C) {
	// Wednesday + Restricted Morning
	lastDigit := 5
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-08T07:30:00-05:00")

	expectedResult := false

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_5(c *C) {
	// Wednesday + Unrestricted Evening
	lastDigit := 6
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-08T15:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_6(c *C) {
	// Thursday + Restricted Morning
	lastDigit := 7
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-09T08:30:00-05:00")

	expectedResult := false

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_7(c *C) {
	// Thursday + Unrestricted Evening
	lastDigit := 8
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-09T15:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_8(c *C) {
	// Friday + Restricted Morning
	lastDigit := 9
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-10T07:30:00-05:00")

	expectedResult := false

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_CheckDigitDifferentDays_9(c *C) {
	// Friday + Unrestricted Evening
	lastDigit := 0
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-10T15:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

// These set of Unit tests, tests for weekend days.
func (s *picoyplacaSuite) Test_allowedInCity_Weekend_1(c *C) {
	lastDigit := 0
	// Saturday
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-11T07:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_allowedInCity_Weekend_2(c *C) {
	lastDigit := 0
	// Sunday
	dateTime, _ := time.Parse(time.RFC3339, "2019-05-12T07:30:00-05:00")

	expectedResult := true

	actualResult := allowedInCity(lastDigit, dateTime)

	c.Assert(actualResult, Equals, expectedResult)
}

func (s *picoyplacaSuite) Test_displayUsage_VerifyText(c *C) {
	os.Args = []string{"/path/to/exec.ext"}

	expectedReturn := 714

	actualResult, actualError := displayUsage()

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}

// Functional tests

func (s *picoyplacaSuite) Test_extractLastDigit_ValidLicensePlate(c *C) {
	validLicensePlate := "AAC-1111"

	expectedLastDigit := 1

	actualLastDigit, actualError := extractLastDigit(validLicensePlate)

	c.Assert(actualError, IsNil)
	c.Assert(actualLastDigit, Equals, expectedLastDigit)
}

func (s *picoyplacaSuite) Test_extractLastDigit_InvalidLicensePlate(c *C) {
	validLicensePlate := "A1C-1111"

	expectedLastDigit := -1

	actualLastDigit, actualError := extractLastDigit(validLicensePlate)

	c.Assert(actualError, NotNil)
	c.Assert(actualError, ErrorMatches, "Invalid license plate")
	c.Assert(actualLastDigit, Equals, expectedLastDigit)
}

func (s *picoyplacaSuite) Test_displayResult_ValidParametersAllowed(c *C) {
	expectedReturn := 46

	actualResult, actualError := displayResult("ABC-1234", "2019-05-20", "09:31:00")

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}

func (s *picoyplacaSuite) Test_displayResult_ValidParametersNotAllowed(c *C) {
	expectedReturn := 50

	actualResult, actualError := displayResult("ABC-1231", "2019-05-20", "07:31:00")

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}

func (s *picoyplacaSuite) Test_dislayResult_InvalidLicense(c *C) {
	expectedReturn := 22

	actualResult, actualError := displayResult("A1C-1234", "2019-05-20", "09:31:00")

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}

func (s *picoyplacaSuite) Test_displayResult_InvalidDate(c *C) {
	expectedReturn := 28

	actualResult, actualError := displayResult("ABC-1234", "2019-15-02", "09:31:00")

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}

func (s *picoyplacaSuite) Test_displayResult_InvalidTime(c *C) {
	expectedReturn := 28

	actualResult, actualError := displayResult("ABC-1234", "2019-15-02", "39:31:99")

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}

func (s *picoyplacaSuite) Test_start_DisplayUsage(c *C) {
	// Usage is displayed when there is not enough parameters
	os.Args = []string{"/path/to/exec.ext", "param1", "param2"}

	expectedReturn := 714

	actualResult, actualError := start()

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}

func (s *picoyplacaSuite) Test_start_DisplayResult(c *C) {
	os.Args = []string{"/path/to/exec.ext", "ABC-1234", "2019-05-20", "09:31:00"}

	expectedReturn := 46

	actualResult, actualError := start()

	c.Assert(actualError, IsNil)
	c.Assert(actualResult, Equals, expectedReturn)
}
