package helpers

import (
	"fmt"

	"github.com/sniddunc/BanLogger/internal/banlogger"
)

// GetInfractionString takes in a StatService implementation and a playerID string and retrieves the user's
// stats and builds and returns an infraction string.
// We take in a StatService instead of attaching this function on an implementation by implementation basis
// because this function relies on any implementation, and requires no special treatment itself.
func GetInfractionString(statService banlogger.StatService, playerID string) string {
	warningCount, _ := statService.GetWarningCount(playerID)
	kickCount, _ := statService.GetKickCount(playerID)
	banCount, _ := statService.GetBanCount(playerID)

	return fmt.Sprintf("Warnings: %d, Kicks: %d, Bans: %d", warningCount, kickCount, banCount)
}
