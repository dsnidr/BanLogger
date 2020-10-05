package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/logging"
)

// MuteService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.MuteService.
type MuteService struct {
	DB *sql.DB
}

// CreateMute inserts a provided mute struct into the sqlite database
func (s *MuteService) CreateMute(mute banlogger.Mute) error {
	query := "INSERT INTO Mute (PlayerID, Duration, Reason, Staff, UnmutedAt, Timestamp) VALUES (?, ?, ?, ?, ?, ?);"

	_, err := s.DB.ExecContext(context.Background(), query,
		mute.PlayerID, mute.Duration, mute.Reason, mute.Staff, mute.UnmutedAt, mute.Timestamp)

	if err != nil {
		return err
	}

	logging.Info("sqlite/mute.go",
		fmt.Sprintf("New mute inserted into the database.\n\tPlayerID: %s | Duration: %s | Reason: %s | Staff: %s | Timestamp: %d",
			mute.PlayerID, mute.Duration, mute.Reason, mute.Staff, mute.Timestamp))

	return nil
}

// GetMutesByPlayerID finds all mutes given to a provided player
func (s *MuteService) GetMutesByPlayerID(playerID string) ([]banlogger.Mute, error) {
	query := "SELECT * FROM Mute WHERE PlayerID = ?;"

	rows, err := s.DB.QueryContext(context.Background(), query, playerID)
	if err != nil {
		return nil, err
	}

	mutes := []banlogger.Mute{}

	for rows.Next() {
		mute := banlogger.Mute{}

		err = rows.Scan(&mute.ID, &mute.PlayerID, &mute.Duration, &mute.Reason, &mute.Staff, &mute.UnmutedAt, &mute.Timestamp)
		if err != nil {
			return nil, err
		}

		mutes = append(mutes, mute)
	}

	return mutes, nil
}
