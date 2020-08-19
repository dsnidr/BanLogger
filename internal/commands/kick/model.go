package kick

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sniddunc/banlogger/pkg/logging"
)

// Kick represents a kick given to a player
type Kick struct {
	ID        int64
	PlayerID  string
	Reason    string
	Staff     string
	Timestamp int64
}

// Insert insters a new kick into the database
func Insert(db *sql.DB, kick Kick) error {
	query := "INSERT INTO Kick (PlayerID, Reason, Staff, Timestamp) VALUES (?, ?, ?, ?);"

	_, err := db.ExecContext(context.Background(), query,
		kick.PlayerID, kick.Reason, kick.Staff, kick.Timestamp)

	if err != nil {
		return err
	}

	logging.Info("kick/models.go",
		fmt.Sprintf("New kick inserted into the database.\n\tPlayerID: %s | Reason: %s | Staff: %s | Timestamp: %d",
			kick.PlayerID, kick.Reason, kick.Staff, kick.Timestamp))

	return nil
}

// FindByPlayerID finds all kicks recorded for a given player
func FindByPlayerID(db *sql.DB, playerID string) ([]Kick, error) {
	query := "SELECT ID, PlayerID, Reason, Staff, Timestamp FROM Kick WHERE PlayerID = ?;"

	rows, err := db.QueryContext(context.Background(), query, playerID)
	if err != nil {
		return nil, err
	}

	kicks := []Kick{}

	for rows.Next() {
		kick := Kick{}

		err = rows.Scan(&kick.ID, &kick.PlayerID, &kick.Reason, &kick.Staff, &kick.Timestamp)
		if err != nil {
			return nil, err
		}

		kicks = append(kicks, kick)
	}

	return kicks, nil
}
