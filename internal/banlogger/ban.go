package banlogger

// Ban represents a ban given to a player
type Ban struct {
	ID         int64
	PlayerID   string
	Duration   string
	Reason     string
	Staff      string
	UnbannedAt int64
	Timestamp  int64
}

// BanService facilitates ban exchange/manipulation between packages
type BanService interface {
	CreateBan(Ban) error
	GetBansByPlayerID(string) ([]Ban, error)
	GetCurrentBans() ([]Ban, error)
}
