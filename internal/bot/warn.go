package bot

import (
	"fmt"
	"strings"

	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/BanLogger/pkg/validation"
	"github.com/sniddunc/gcmd"
)

// ParseWarnArgs parses the received arguments for the warn command. It's main function is to
// detect errors, report these errors and if no errors occur, then it attaches the necessary data
// to the provided command context.
func (bot *Bot) ParseWarnArgs(next gcmd.HandlerFunc) gcmd.HandlerFunc {
	return func(c gcmd.Context) error {
		args := c.Args

		// Check amount of arguments provided
		if len(args) < 2 {
			return fmt.Errorf("Invalid arguments")
		}

		profileURL := args[0]

		// Validate profileURL
		ok, isVanity := validation.ValidateProfileURL(profileURL)
		if !ok {
			// If it wasn't a vanity or profile URL, check if it's just a SteamID64

			return fmt.Errorf("Invalid profile URL provided")
		}

		reason := strings.Join(args[1:], " ")

		if len(reason) < config.ReasonMinLen || len(reason) > config.ReasonMaxLen {
			return fmt.Errorf("Reason must be between %d and %d in length", config.ReasonMinLen, config.ReasonMaxLen)
		}

		// Get the user's SteamID
		steamID, err := bot.SteamService.GetSteamID(isVanity, profileURL)
		if err != nil {
			return fmt.Errorf("Could not resolve SteamID")
		}

		// Map data to context store
		c.Set("reason", reason)
		c.Set("steamID", steamID)

		logging.Info("bot/warn.go", fmt.Sprintf("Warn command passed arguments check. steamID: %s | reason: %s", steamID, reason))

		return next(c)
	}
}
