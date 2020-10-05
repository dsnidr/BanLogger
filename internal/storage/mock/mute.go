package mock

import (
	"github.com/sniddunc/BanLogger/internal/banlogger"
)

var currentMuteID int64 = 0

// MuteService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.MuteService.
type MuteService struct {
	Mutes []banlogger.Mute
}

// CreateMute storages a ban record
func (s *MuteService) CreateMute(mute banlogger.Mute) error {
	mute.ID = currentMuteID
	currentMuteID++

	s.Mutes = append(s.Mutes, mute)

	return nil
}

// GetMutesByPlayerID finds all mutes given to a provided player
func (s *MuteService) GetMutesByPlayerID(playerID string) ([]banlogger.Mute, error) {
	mutes := []banlogger.Mute{}

	for _, mute := range s.Mutes {
		if mute.PlayerID == playerID {
			mutes = append(mutes, mute)
		}
	}

	return mutes, nil
}
