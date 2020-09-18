package mock

import (
	"github.com/sniddunc/BanLogger/internal/banlogger"
)

var currentWarningID int64 = 0

// WarningService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.WarningService.
type WarningService struct {
	Warnings []banlogger.Warning
}

// CreateWarning records a warning
func (s *WarningService) CreateWarning(warning banlogger.Warning) error {
	warning.ID = currentWarningID
	currentWarningID++

	s.Warnings = append(s.Warnings, warning)

	return nil
}

// GetWarningsByPlayerID finds all warnings given to a provided player
func (s *WarningService) GetWarningsByPlayerID(playerID string) ([]banlogger.Warning, error) {
	warnings := []banlogger.Warning{}

	for _, warning := range s.Warnings {
		if warning.PlayerID == playerID {
			warnings = append(warnings, warning)
		}
	}

	return warnings, nil
}
