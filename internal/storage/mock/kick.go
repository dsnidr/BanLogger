package mock

import "github.com/sniddunc/BanLogger/internal/banlogger"

var currentKickID int64 = 0

// KickService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.KickService.
type KickService struct {
	Kicks []banlogger.Kick
}

// CreateKick records a kick
func (s *KickService) CreateKick(kick banlogger.Kick) error {
	kick.ID = currentKickID
	currentKickID++

	s.Kicks = append(s.Kicks, kick)

	return nil
}

// GetKicksByPlayerID finds all kicks given to a provided player
func (s *KickService) GetKicksByPlayerID(playerID string) ([]banlogger.Kick, error) {
	kicks := []banlogger.Kick{}

	for _, kick := range s.Kicks {
		if kick.PlayerID == playerID {
			kicks = append(kicks, kick)
		}
	}

	return kicks, nil
}
