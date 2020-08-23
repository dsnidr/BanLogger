package banlogger

// Warning represents a warning given to a player
type Warning struct {
	ID        int64
	PlayerID  string
	Reason    string
	Staff     string
	Timestamp int64
}

// WarningService facilitates warning exchange/manipulation between packages
type WarningService interface {
	CreateWarning(Warning) error
	GetWarningsByPlayerID(string) ([]Warning, error)
}
