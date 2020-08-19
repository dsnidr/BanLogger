package lookup

import (
	"fmt"

	"github.com/sniddunc/banlogger/internal/steam"
	"github.com/sniddunc/banlogger/pkg/logging"
	"github.com/sniddunc/banlogger/pkg/validation"
	"github.com/sniddunc/gcmd"
)

// ValidateAndMapArgs validates the provided arguments and maps them out to the store on context
// so that they can be easily used by the command handler.
func ValidateAndMapArgs(next gcmd.HandlerFunc) gcmd.HandlerFunc {
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
		steamID, err := steam.GetSteamID(isVanity, profileURL)
		if err != nil {
			return fmt.Errorf("Could not resolve SteamID")
		}

		// Map data to context store
		c.Set("steamID", steamID)

		logging.Info("lookup/args.go", fmt.Sprintf("Lookup command passed arguments check. steamID: %s", steamID))

		return next(c)
	}
}
