package stats

import "database/sql"

// GetWarnCount returns the number of warning records this user has recorded
func GetWarnCount(db *sql.DB, playerID string) (int, error) {
	query := "SELECT COUNT(1) AS Count FROM Warning WHERE PlayerID = ?;"

	row := db.QueryRow(query, playerID)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetKickCount returns the number of kick records this user has recorded
func GetKickCount(db *sql.DB, playerID string) (int, error) {
	query := "SELECT COUNT(1) AS Count FROM Kick WHERE PlayerID = ?;"

	row := db.QueryRow(query, playerID)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetBanCount returns the number of ban records this user has recorded
func GetBanCount(db *sql.DB, playerID string) (int, error) {
	query := "SELECT COUNT(1) AS Count FROM Ban WHERE PlayerID = ?;"

	row := db.QueryRow(query, playerID)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
