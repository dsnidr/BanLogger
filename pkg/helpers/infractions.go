package helpers

import (
	"database/sql"
	"fmt"

	"github.com/sniddunc/banlogger/internal/stats"
)

// GetInfractionString returns a string informing the user of how many infractions someone has
func GetInfractionString(db *sql.DB, playerID string) string {
	warningCount, _ := stats.GetWarnCount(db, playerID)
	kickCount, _ := stats.GetKickCount(db, playerID)
	banCount, _ := stats.GetBanCount(db, playerID)

	return fmt.Sprintf("Warnings: %d, Kicks: %d, Bans: %d", warningCount, kickCount, banCount)
}
