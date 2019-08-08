package main

import (
	"errors"
	"strings"
	"time"
	"unicode"
)

func main() {
	// XXX: Not Implemented
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
func parseDateTime(dateString, timeString string) (time.Time, error) {
	dateTime, err := time.Parse(time.RFC3339, dateString + "T" + timeString + "-05:00")
	if err != nil {
		return dateTime, errors.New("Invalid Date or Time given.")
	}
	return dateTime, err
}

// Extracts the last digit of the given license plate
func extractLastDigit(licensePlate string) (int, error) {
	// XXX: Not Implemented
	return 0, nil
}

// Applies the rules of "Pico y Placa" to see if the vechicle is allowed in the
// city.
// Reference: https://es.wikipedia.org/wiki/Pico_y_placa#Quito,_Ecuador
func allowedInCity(lastDigit int, dateTime time.Time) bool {
	// XXX: Not Implemented
	return false
}
