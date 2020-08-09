package lookup

import (
	"database/sql"
	"fmt"

	"github.com/sniddunc/banlogger/internal/ban"
	"github.com/sniddunc/banlogger/internal/kick"
	"github.com/sniddunc/banlogger/internal/warn"
	"github.com/sniddunc/banlogger/pkg/logging"
)

// Record represents a player's punitive record
type Record struct {
	PlayerID string
	Warnings []warn.Warning
	Kicks    []kick.Kick
	Bans     []ban.Ban
}

// GetRecord builds a player's record from their warnings, kicks and bans
func GetRecord(db *sql.DB, playerID string) (Record, error) {
	warnings, err := warn.FindByPlayerID(db, playerID)
	if err != nil {
		return Record{}, err
	}

	kicks, err := kick.FindByPlayerID(db, playerID)
	if err != nil {
		return Record{}, err
	}

	bans, err := ban.FindByPlayerID(db, playerID)
	if err != nil {
		return Record{}, err
	}

	record := Record{
		PlayerID: playerID,
		Warnings: warnings,
		Kicks:    kicks,
		Bans:     bans,
	}

	logging.Info("lookup/models.go",
		fmt.Sprintf("Record retrieved for player %s.\n\tWarnings: %d | Kicks: %d | Bans: %d",
			playerID, len(record.Warnings), len(record.Kicks), len(record.Bans)))

	return record, nil
}
