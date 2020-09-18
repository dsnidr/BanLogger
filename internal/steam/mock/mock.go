package live

import (
	"log"
	"math/rand"
	"strings"

	"github.com/sniddunc/BanLogger/internal/banlogger"
)

// SteamService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.SteamService.
type SteamService struct{}

// GetSteamID retrieves a player's SteamID64 from either their profile URL, or their vanity URL
func (s *SteamService) GetSteamID(resolveVanity bool, url string) (string, error) {
	// Remove trailing slash if one is present
	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	// Grab the ID portion of the url
	split := strings.Split(url, "/")
	idPart := split[len(split)-1]

	steamID := idPart

	if resolveVanity {
		log.Println("Can't resolve vanity in mock mode. Returning a valid fake ID")

		return randomSteamID(), nil
	}

	return steamID, nil
}

// GetUserSummary takes in a user's steam ID and gets their profile summary
func (s *SteamService) GetUserSummary(steamID string) (banlogger.SteamPlayerSummary, error) {
	return banlogger.SteamPlayerSummary{
		AvatarURL:   "https://homepages.cae.wisc.edu/~ece533/images/airplane.png",
		ProfileName: "Test",
		ProfileURL:  "https://steamcommunity.com/id/tdfubghwgwhdghwcgdubhghuwguygd",
	}, nil
}

func randomSteamID() string {
	const letters = "123456789"

	bytes := make([]byte, 17)

	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}
