package steam

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type steamVanityResponseWrapper struct {
	Response steamVanityResponse `json:"response"`
}

type steamVanityResponse struct {
	SteamID string `json:"steamid"`
}

// GetSteamID retrieves a player's SteamID64 from either their profile URL, or their vanity URL
func GetSteamID(resolveVanity bool, url string) (string, error) {
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

		var result steamVanityResponseWrapper
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			log.Println(err)
			return "", fmt.Errorf("Could not decode API response")
		}

		steamID = result.Response.SteamID
	}

	return steamID, nil
}
