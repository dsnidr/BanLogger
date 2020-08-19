package ban

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sniddunc/banlogger/pkg/logging"
)

// Ban represents a ban given to a player
type Ban struct {
	ID        int64
	PlayerID  string
	Duration  string
	Reason    string
	Staff     string
	Timestamp int64
}

// Insert inserts a new ban into the database
func Insert(db *sql.DB, ban Ban) error {
	query := "INSERT INTO Ban (PlayerID, Duration, Reason, Staff, Timestamp) VALUES (?, ?, ?, ?, ?);"

	_, err := db.ExecContext(context.Background(), query,
		ban.PlayerID, ban.Duration, ban.Reason, ban.Staff, ban.Timestamp)

	if err != nil {
		return err
	}

	logging.Info("ban/models.go",
		fmt.Sprintf("New ban inserted into the database.\n\tPlayerID: %s | Duration: %s | Reason: %s | Staff: %s | Timestamp: %d",
			ban.PlayerID, ban.Duration, ban.Reason, ban.Staff, ban.Timestamp))

	return nil
}

// FindByPlayerID finds all bans records for a given player
func FindByPlayerID(db *sql.DB, playerID string) ([]Ban, error) {
	query := "SELECT ID, PlayerID, Duration, Reason, Staff, Timestamp FROM Ban WHERE PlayerID = ?;"

	rows, err := db.QueryContext(context.Background(), query, playerID)
	if err != nil {
		return nil, err
	}

	bans := []Ban{}

	for rows.Next() {
		ban := Ban{}

		err = rows.Scan(&ban.ID, &ban.PlayerID, &ban.Duration, &ban.Reason, &ban.Staff, &ban.Timestamp)
		if err != nil {
			return nil, err
		}

		bans = append(bans, ban)
	}

	return bans, nil
}
