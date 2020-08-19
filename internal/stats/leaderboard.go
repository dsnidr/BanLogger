package stats

import (
	"database/sql"
)

// PlayerInfractionRecord holds infraction stats
type PlayerInfractionRecord struct {
	PlayerID         string
	TotalInfractions int
	WarnCount        int
	KickCount        int
	BanCount         int
}

// GetPlayerWithMostInfractions retrieves the player with the highest number of infractions
func GetPlayerWithMostInfractions(db *sql.DB) (PlayerInfractionRecord, error) {
	query := `
	SELECT * FROM (
		SELECT
			w.PlayerID
		FROM Warning w
		INNER JOIN Kick k ON w.PlayerID = k.PlayerID 
		INNER JOIN Ban b ON w.PlayerID = b.PlayerID AND k.PlayerID = b.PlayerID
	) common
	GROUP BY common.PlayerID ORDER BY COUNT(PlayerID) DESC LIMIT 1
`

	row := db.QueryRow(query)

	var record PlayerInfractionRecord

	err := row.Scan(&record.PlayerID)
	if err != nil {
		return PlayerInfractionRecord{}, err
	}

	warningCount, err := GetWarnCount(db, record.PlayerID)
	if err != nil {
		return PlayerInfractionRecord{}, err
	}

	kickCount, err := GetKickCount(db, record.PlayerID)
	if err != nil {
		return PlayerInfractionRecord{}, err
	}

	banCount, err := GetBanCount(db, record.PlayerID)
	if err != nil {
		return PlayerInfractionRecord{}, err
	}

	record.TotalInfractions = warningCount + kickCount + banCount
	record.WarnCount = warningCount
	record.KickCount = kickCount
	record.BanCount = banCount

	return record, err
}
