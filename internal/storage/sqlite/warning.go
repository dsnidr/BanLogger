package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/logging"
)

// WarningService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.WarningService.
type WarningService struct {
	DB *sql.DB
}

// CreateWarning inserts a provided warning struct into the sqlite database
func (s *WarningService) CreateWarning(warning banlogger.Warning) error {
	query := "INSERT INTO Warning (PlayerID, Reason, Staff, Timestamp) VALUES (?, ?, ?, ?);"

	_, err := s.DB.ExecContext(context.Background(), query,
		warning.PlayerID, warning.Reason, warning.Staff, warning.Timestamp)

	if err != nil {
		return err
	}

	logging.Info("sqlite/warning.go",
		fmt.Sprintf("New warning inserted into the database.\n\tPlayerID: %s | Reason: %s | Staff: %s | Timestamp: %d",
			warning.PlayerID, warning.Reason, warning.Staff, warning.Timestamp))

	return nil
}

// GetWarningsByPlayerID finds all warnings given to a provided player
func (s *WarningService) GetWarningsByPlayerID(playerID string) ([]banlogger.Warning, error) {
	query := "SELECT ID, PlayerID, Reason, Staff, Timestamp FROM Warning WHERE PlayerID = ?;"

	rows, err := s.DB.QueryContext(context.Background(), query, playerID)
	if err != nil {
		return nil, err
	}

	warnings := []banlogger.Warning{}

	for rows.Next() {
		warning := banlogger.Warning{}

		err = rows.Scan(&warning.ID, &warning.PlayerID, &warning.Reason, &warning.Staff, &warning.Timestamp)
		if err != nil {
			return nil, err
		}

		warnings = append(warnings, warning)
	}

	return warnings, nil
}
