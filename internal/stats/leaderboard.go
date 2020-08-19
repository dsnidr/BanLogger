package stats

import "database/sql"

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
	SELECT 
		top.*,
		wc.WarningCount,
		kc.KickCount,
		bc.BanCount
	FROM (
		SELECT
			w.PlayerID,
			COUNT(w.PlayerID) as InfractionCount
		FROM Warning w
		INNER JOIN Kick k ON k.PlayerID = w.PlayerID
		INNER JOIN Ban b ON b.PlayerID = w.PlayerID
	) top
	INNER JOIN (
		SELECT PlayerID, COUNT(1) AS WarningCount FROM Warning
	) wc ON wc.PlayerID = top.PlayerID
	INNER JOIN (
		SELECT PlayerID, COUNT(1) AS KickCount FROM Kick
	) kc ON kc.PlayerID = top.PlayerID
	INNER JOIN (
		SELECT PlayerID, COUNT(1) AS BanCount FROM Kick
	) bc ON bc.PlayerID = top.PlayerID
	`

	row := db.QueryRow(query)

	var record PlayerInfractionRecord

	err := row.Scan(&record.PlayerID, &record.TotalInfractions, &record.WarnCount, &record.KickCount, &record.BanCount)
	if err != nil {
		return PlayerInfractionRecord{}, err
	}
	return record, err
}
