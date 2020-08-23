package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/logging"
)

// KickService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.KickService.
type KickService struct {
	DB *sql.DB
}

// CreateKick inserts a provided kick struct into the sqlite database
func (s *KickService) CreateKick(kick banlogger.Kick) error {
	query := "INSERT INTO Kick (PlayerID, Reason, Staff, Timestamp) VALUES (?, ?, ?, ?);"

	_, err := s.DB.ExecContext(context.Background(), query,
		kick.PlayerID, kick.Reason, kick.Staff, kick.Timestamp)

	if err != nil {
		return err
	}

	logging.Info("sqlite/kick.go",
		fmt.Sprintf("New kick inserted into the database.\n\tPlayerID: %s | Reason: %s | Staff: %s | Timestamp: %d",
			kick.PlayerID, kick.Reason, kick.Staff, kick.Timestamp))

	return nil
}

// GetKicksByPlayerID finds all kicks given to a provided player
func (s *KickService) GetKicksByPlayerID(playerID string) ([]banlogger.Kick, error) {
	query := "SELECT ID, PlayerID, Reason, Staff, Timestamp FROM Kick WHERE PlayerID = ?;"

	rows, err := s.DB.QueryContext(context.Background(), query, playerID)
	if err != nil {
		return nil, err
	}

	kicks := []banlogger.Kick{}

	for rows.Next() {
		kick := banlogger.Kick{}

		err = rows.Scan(&kick.ID, &kick.PlayerID, &kick.Reason, &kick.Staff, &kick.Timestamp)
		if err != nil {
			return nil, err
		}

		kicks = append(kicks, kick)
	}

	return kicks, nil
}
