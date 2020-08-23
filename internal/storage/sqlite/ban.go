package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/logging"
)

// BanService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.BanService.
type BanService struct {
	DB *sql.DB
}

// CreateBan inserts a provided ban struct into the sqlite database
func (s *BanService) CreateBan(ban banlogger.Ban) error {
	query := "INSERT INTO Ban (PlayerID, Duration, Reason, Staff, Timestamp) VALUES (?, ?, ?, ?, ?);"

	_, err := s.DB.ExecContext(context.Background(), query,
		ban.PlayerID, ban.Duration, ban.Reason, ban.Staff, ban.Timestamp)

	if err != nil {
		return err
	}

	logging.Info("sqlite/ban.go",
		fmt.Sprintf("New ban inserted into the database.\n\tPlayerID: %s | Duration: %s | Reason: %s | Staff: %s | Timestamp: %d",
			ban.PlayerID, ban.Duration, ban.Reason, ban.Staff, ban.Timestamp))

	return nil
}

// GetBansByPlayerID finds all bans given to a provided player
func (s *BanService) GetBansByPlayerID(playerID string) ([]banlogger.Ban, error) {
	query := "SELECT ID, PlayerID, Duration, Reason, Staff, Timestamp FROM Ban WHERE PlayerID = ?;"

	rows, err := s.DB.QueryContext(context.Background(), query, playerID)
	if err != nil {
		return nil, err
	}

	bans := []banlogger.Ban{}

	for rows.Next() {
		ban := banlogger.Ban{}

		err = rows.Scan(&ban.ID, &ban.PlayerID, &ban.Duration, &ban.Reason, &ban.Staff, &ban.Timestamp)
		if err != nil {
			return nil, err
		}

		bans = append(bans, ban)
	}

	return bans, nil
}
