package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	// XXX: Not Implemented
}

func start() (int, error) {
	if len(os.Args) != 4 {
		return displayUsage()
	}

	licensePlate := os.Args[1]
	dateString := os.Args[2]
	timeString := os.Args[3]

	return displayResult(licensePlate, dateString, timeString)
}

func displayUsage() (int, error) {
	executableName := strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(filepath.Base(os.Args[0])))
	helpText := "Usage: " + executableName  + " LICENSE-PLATE DATE TIME\n" +
		"Check if a LICENSE-PLATE is allowed within Quito City for a given DATE and TIME,\n"+
		"Pico y Placa policy implemented in Quito.\n" +
		"Example: " + executableName + " AAB-0123 2019-05-20 09:31:00\n\n"+
		"Interpretation of parameters:\n"+
		"\tLICENSE-PLATE\tThis is the license plate number as described in the legal\n"+
		"\t\t\tsystem of Ecuador, only Standard Plates are implemented as of now.\n"+
		"\t\t\tFor example: ABC-0123\n\n"+

		"\tDATE\t\tThe date to check for if a vehicle is allowed in the city.\n"+
		"\t\t\tThe input expects ISO 8601 standard of the form YYYY-MM-DD.\n"+
		"\t\t\tFor example: 2019-05-20\n\n"+

		"\tTIME\t\tThe time to check for if a vehicle is allowed in the city.\n"+
		"\t\t\tThe input expects ISO 8601 standard of the form HH:MM:SS.\n"+
		"\t\t\tFor example: 09:31:00\n\n"

	bytesWritten, err := fmt.Println(helpText)

	return bytesWritten, err
}

func displayResult(licensePlate, dateString, timeString string) (int, error) {
	var bytesWritten int
	var err error

	lastDigit, errLicense := extractLastDigit(licensePlate)

	if errLicense == nil {
		dateTime, errDateTime := parseDateTime(dateString, timeString)
		if errDateTime == nil {
			if allowedInCity(lastDigit, dateTime) {
				bytesWritten, err = fmt.Println(licensePlate + " is allowed on " + dateString + " at " + timeString)
			} else {
				bytesWritten, err = fmt.Println(licensePlate + " is NOT allowed on " + dateString + " at " + timeString)
			}
		} else {
			bytesWritten, err = fmt.Println(errDateTime)
		}
	} else {
		bytesWritten, err = fmt.Println(errLicense)
	}
	return bytesWritten, err
}

// A custom function to validate a given license plate value.
//
// TODO: One can improve upon this function to add additional checks
//
// 1. Better checking of plates to identify if they are government, diplomatic,
// offical etc, this can be better used to filter out vehicles to which the
// Pico y Placa rules do not apply
//
// 2. Add some way to identify emergency vehicles
//
// Reference: https://en.wikipedia.org/wiki/Vehicle_registration_plates_of_Ecuador
func validateLicensePlate(licensePlate string) bool {
	licensePlateSplit := strings.Split(licensePlate, "-")

	// There should be only two parts when split by hyphen
	if len(licensePlateSplit) != 2 {
		return false
	}

	// Validate first part is purely alphabets
	for _, r := range licensePlateSplit[0] {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	// Validate second part is purely digits
	for _, r := range licensePlateSplit[1] {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// Converts an ISO 8601 notation date (YYYY-MM-DD) and time (HH:MM:SS) into a
// RFC3339 Time object for further processing.
// NOTE: We do UTC-05:00 to conform it to Quito time.
func parseDateTime(dateString, timeString string) (time.Time, error) {
	dateTime, err := time.Parse(time.RFC3339, dateString + "T" + timeString + "-05:00")
	if err != nil {
		return dateTime, errors.New("Invalid Date or Time given.")
	}
	return dateTime, err
}

// Extracts the last digit of the given license plate.
// In case of invalid plates we return a -1, since license plates numbers cannot
// be negative.
func extractLastDigit(licensePlate string) (int, error) {
	if validateLicensePlate(licensePlate) == true {
		lastDigit, _ := strconv.Atoi(licensePlate[len(licensePlate) - 1:])
		return lastDigit, nil
	} else {
		return -1, errors.New("Invalid license plate")
	}
}

// Applies the rules of "Pico y Placa" to see if the vechicle is allowed in the
// city.
// NOTE: Holidays are not considered in this logic.
// Reference: https://es.wikipedia.org/wiki/Pico_y_placa#Quito,_Ecuador
func allowedInCity(lastDigit int, dateTime time.Time) bool {
	allowed := true

	weekday := dateTime.Weekday()
	dateString := dateTime.Format("2006-01-02")

	morningRestrictionStart, _ := time.Parse(time.RFC3339, dateString + "T" + "07:00:00-05:00")
	morningRestrictionEnd, _ := time.Parse(time.RFC3339, dateString + "T" + "09:30:00-05:00")

	eveningRestrictionStart, _ := time.Parse(time.RFC3339, dateString + "T" + "16:00:00-05:00")
	eveningRestrictionEnd, _ := time.Parse(time.RFC3339, dateString + "T" + "19:30:00-05:00")

	restrictedMorningTime := dateTime.After(morningRestrictionStart) && dateTime.Before(morningRestrictionEnd)
	restrictedEveningTime := dateTime.After(eveningRestrictionStart) && dateTime.Before(eveningRestrictionEnd)

	if restrictedMorningTime || restrictedEveningTime {
		switch weekday {
		case time.Monday:
			if lastDigit == 1 || lastDigit == 2 {
				allowed = false
			}
		case time.Tuesday:
			if lastDigit == 3 || lastDigit == 4 {
				allowed = false
			}
		case time.Wednesday:
			if lastDigit == 5 || lastDigit == 6 {
				allowed = false
			}
		case time.Thursday:
			if lastDigit == 7 || lastDigit == 8 {
				allowed = false
			}
		case time.Friday:
			if lastDigit == 9 || lastDigit == 0 {
				allowed = false
			}
		}
	}

	return allowed
}
