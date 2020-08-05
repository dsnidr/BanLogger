package validation

import "regexp"

const (
	steamIDRegex = "^[0-9]{17}$"
)

// ValidateSteamID takes in a string and returns true if it's a valid SteamID64.
func ValidateSteamID(steamID string) bool {
	match, err := regexp.MatchString(steamIDRegex, steamID)
	if err != nil || !match {
		return false
	}

	return true
}
