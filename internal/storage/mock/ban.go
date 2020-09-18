package mock

import (
	"github.com/sniddunc/BanLogger/internal/banlogger"
)

var currentBanID int64 = 0

// BanService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.BanService.
type BanService struct {
	Bans []banlogger.Ban
}

// CreateBan storages a ban record
func (s *BanService) CreateBan(ban banlogger.Ban) error {
	ban.ID = currentBanID
	currentBanID++

	s.Bans = append(s.Bans, ban)

	return nil
}

// GetBansByPlayerID finds all bans given to a provided player
func (s *BanService) GetBansByPlayerID(playerID string) ([]banlogger.Ban, error) {
	bans := []banlogger.Ban{}

	for _, ban := range s.Bans {
		if ban.PlayerID == playerID {
			bans = append(bans, ban)
		}
	}

	return bans, nil
}
