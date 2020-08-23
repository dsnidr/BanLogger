package banlogger

// Kick represents a kick given to a player
type Kick struct {
	ID        int64
	PlayerID  string
	Reason    string
	Staff     string
	Timestamp int64
}

// KickService facilitates kick exchange/manipulation between packages
type KickService interface {
	CreateKick(Kick) error
	GetKicksByPlayerID(string) ([]Kick, error)
}
