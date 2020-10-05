package banlogger

// Mute represents a mute given to a player
type Mute struct {
	ID        int64
	PlayerID  string
	Duration  string
	Reason    string
	Staff     string
	UnmutedAt int64
	Timestamp int64
}

// MuteService facilitates kick exchange/manipulation between packages
type MuteService interface {
	CreateMute(Mute) error
	GetMutesByPlayerID(string) ([]Mute, error)
}
