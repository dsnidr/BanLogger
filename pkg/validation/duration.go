package validation

import (
	"regexp"
)

const (
	durationRegex = "^([0-9]{1,5}([hdmy]|min))$|^perm$"
)

// ValidateDuration takes in a string and verifies that it's in the expected format.
// It returns true if it does, and false otherwise.
func ValidateDuration(duration string) bool {
	match, err := regexp.MatchString(durationRegex, duration)
	if err != nil || !match {
		return false
	}

	return true
}
