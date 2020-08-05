package validation

import (
	"regexp"
)

const (
	vanityURLRegex  = "^https:\\/\\/steamcommunity\\.com\\/id\\/[a-zA-Z0-9_]+\\/?$"
	profileURLRegex = "^https:\\/\\/steamcommunity\\.com\\/profiles\\/[0-9]{17,17}\\/?$"
)

// ValidateProfileURL returns two booleans. The first one is true if it's a valid profile URL
// and the second one is set to true if the tested URL is a vanity URL that needs to be resolved
func ValidateProfileURL(profileURL string) (bool, bool) {
	// Check if URL is vanity URL
	match, err := regexp.MatchString(vanityURLRegex, profileURL)
	if err != nil {
		return false, false
	}

	// If it was a vanity URL, set a flag and continue
	if match {
		return true, true
	}

	// If it's not a vanity URL, check if it's a profile URL
	match, err = regexp.MatchString(profileURLRegex, profileURL)
	if err != nil {
		return false, false
	}

	// If it was a profile URL, set a flag and continue
	if match {
		return true, false
	}

	// Otherwise if it wasn't any of the above, return an error
	return false, false
}
