package banlogger

// SteamVanityResponseWrapper represents the top level of the response we get when
// resolving a vanity URL
type SteamVanityResponseWrapper struct {
	Response SteamVanityResponse `json:"response"`
}

// SteamVanityResponse represents the mid level of the response we get when
// resolving a vanity URL
type SteamVanityResponse struct {
	SteamID string `json:"steamid"`
}

// SteamSummaryResponseWrapper represents the top level of the response we get when
// we fetch a player's summary
type SteamSummaryResponseWrapper struct {
	Response SteamSummaryResponse `json:"response"`
}

// SteamSummaryResponse represents the mid level of the response we get when
// we fetch a player's summary
type SteamSummaryResponse struct {
	Players []SteamPlayerSummary `json:"players"`
}

// SteamPlayerSummary represents data retrieved from a summary call to steam's API
type SteamPlayerSummary struct {
	ProfileName string `json:"personaname"`
	ProfileURL  string `json:"profileurl"`
	AvatarURL   string `json:"avatarmedium"`
}

// SteamService facilitates steam data exchange/manipulation between packages
type SteamService interface {
	GetSteamID(bool, string) (string, error)
	GetUserSummary(string) (SteamPlayerSummary, error)
}
