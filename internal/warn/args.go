package warn

import (
	"fmt"
	"strings"

	"github.com/sniddunc/banlogger/internal/steam"
	"github.com/sniddunc/banlogger/pkg/validation"
	"github.com/sniddunc/gcmd"
)

const (
	reasonMinLen = 1
	reasonMaxLen = 128
)

// ValidateAndMapArgs validates the provided arguments and maps them out to the store on context
// so that they can be easily used by the command handler.
func ValidateAndMapArgs(next gcmd.HandlerFunc) gcmd.HandlerFunc {
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

		if len(reason) < reasonMinLen || len(reason) > reasonMaxLen {
			return fmt.Errorf("Reason must be between %d and %d in length", reasonMinLen, reasonMaxLen)
		}

		// Get the user's SteamID
		steamID, err := steam.GetSteamID(isVanity, profileURL)
		if err != nil {
			return fmt.Errorf("Could not resolve SteamID")
		}

		// Map data to context store
		c.Set("reason", reason)
		c.Set("steamID", steamID)

		return next(c)
	}
}
