package warn

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sniddunc/banlogger/pkg/logging"
)

// Warning represents a warning given to a player
type Warning struct {
	ID        int64
	PlayerID  string
	Reason    string
	Staff     string
	Timestamp int64
}

// Insert insters a new warning into the database
func Insert(db *sql.DB, warning Warning) error {
	query := "INSERT INTO Warning (PlayerID, Reason, Staff, Timestamp) VALUES (?, ?, ?, ?);"

	_, err := db.ExecContext(context.Background(), query,
		warning.PlayerID, warning.Reason, warning.Staff, warning.Timestamp)

	if err != nil {
		return err
	}

	logging.Info("warning/models.go",
		fmt.Sprintf("New warning inserted into the database.\n\tPlayerID: %s | Reason: %s | Staff: %s | Timestamp: %d",
			warning.PlayerID, warning.Reason, warning.Staff, warning.Timestamp))

	return nil
}

// FindByPlayerID finds all warn records for a given player
func FindByPlayerID(db *sql.DB, playerID string) ([]Warning, error) {
	query := "SELECT ID, PlayerID, Reason, Staff, Timestamp FROM Warning WHERE PlayerID = ?;"

	rows, err := db.QueryContext(context.Background(), query, playerID)
	if err != nil {
		return nil, err
	}

	warnings := []Warning{}

	for rows.Next() {
		warning := Warning{}

		err = rows.Scan(&warning.ID, &warning.PlayerID, &warning.Reason, &warning.Staff, &warning.Timestamp)
		if err != nil {
			return nil, err
		}

		warnings = append(warnings, warning)
	}

	return warnings, nil
}
