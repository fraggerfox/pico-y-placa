package main

import (
	"time"
)

func main() {
	// XXX: Not Implemented
}

// A custom function to validate a given license plate value.
// Reference: https://en.wikipedia.org/wiki/Vehicle_registration_plates_of_Ecuador
func validateLicensePlate(licensePlate string) bool {
	// XXX: Not Implemented
	return false
}

// Converts an ISO 8601 notation date (YYYY-MM-DD) and time (HH:MM:SS) into a
// RFC3339 Time object for further processing.
func parseDateTime(dateString, timeString string) (time.Time, error) {
	// XXX: Not Implemented
	return time.Now(), nil
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
