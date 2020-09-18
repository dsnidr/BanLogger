package bot

import (
	"fmt"

	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/BanLogger/pkg/validation"
	"github.com/sniddunc/gcmd"
)

// ParseLookupArgs parses the received arguments for the lookup command. It's main function is to
// detect errors, report these errors and if no errors occur, then it attaches the necessary data
// to the provided command context.
func (bot *Bot) ParseLookupArgs(next gcmd.HandlerFunc) gcmd.HandlerFunc {
	return func(c gcmd.Context) error {
		args := c.Args

		// Check amount of arguments provided
		if len(args) < 1 {
			return fmt.Errorf("Invalid arguments")
		}

		profileURL := args[0]

		// Validate profileURL
		ok, isVanity := validation.ValidateProfileURL(profileURL)
		if !ok {
			return fmt.Errorf("Invalid profile URL provided")
		}

		// Get the user's SteamID
		steamID, err := bot.SteamService.GetSteamID(isVanity, profileURL)
		if err != nil {
			return fmt.Errorf("Could not resolve SteamID")
		}

		// Map data to context store
		c.Set("steamID", steamID)

		logging.Info("bot/lookup.go", fmt.Sprintf("Lookup command passed arguments check. steamID: %s", steamID))

		return next(c)
	}
}
