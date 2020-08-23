package live

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
		reqURL := fmt.Sprintf("https://api.steampowered.com/ISteamUser/ResolveVanityURL/v1/?key=%s&vanityurl=%s", os.Getenv("STEAM_API_KEY"), idPart)

		res, err := http.Get(reqURL)
		if err != nil {
			return "", fmt.Errorf("Could not resolve vanity URL. Please try again later")
		}

		var result banlogger.SteamVanityResponseWrapper
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			log.Println(err)
			return "", fmt.Errorf("Could not decode API response")
		}

		steamID = result.Response.SteamID
	}

	return steamID, nil
}

// GetUserSummary takes in a user's steam ID and gets their profile summary
func (s *SteamService) GetUserSummary(steamID string) (banlogger.SteamPlayerSummary, error) {
	reqURL := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", os.Getenv("STEAM_API_KEY"), steamID)

	res, err := http.Get(reqURL)
	if err != nil {
		return banlogger.SteamPlayerSummary{}, fmt.Errorf("Could not get profile summary")
	}

	var result banlogger.SteamSummaryResponseWrapper
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return banlogger.SteamPlayerSummary{}, fmt.Errorf("Could not decode API response")
	}

	summaries := result.Response.Players

	if len(summaries) == 0 {
		return banlogger.SteamPlayerSummary{}, fmt.Errorf("Could not get profile summary")
	}

	return summaries[0], nil
}
