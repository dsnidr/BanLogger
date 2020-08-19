package stats

import (
	"database/sql"
	"fmt"
)

// GetWarnCount returns the number of warning records this user has recorded
func GetWarnCount(db *sql.DB, playerID string) (int, error) {
	return getCountForPlayer(db, "Warning", playerID)
}

// GetKickCount returns the number of kick records this user has recorded
func GetKickCount(db *sql.DB, playerID string) (int, error) {
	return getCountForPlayer(db, "Kick", playerID)
}

// GetBanCount returns the number of ban records this user has recorded
func GetBanCount(db *sql.DB, playerID string) (int, error) {
	return getCountForPlayer(db, "Ban", playerID)
}

// GetTotalWarnCount returns the total number of warning records for any user
func GetTotalWarnCount(db *sql.DB) (int, error) {
	return getTotalCount(db, "Warning")
}

// GetTotalKickCount returns the total number of kick records for any user
func GetTotalKickCount(db *sql.DB) (int, error) {
	return getTotalCount(db, "Kick")
}

// GetTotalBanCount returns the total number of ban records for any user
func GetTotalBanCount(db *sql.DB) (int, error) {
	return getTotalCount(db, "Ban")
}

func getCountForPlayer(db *sql.DB, table string, playerID string) (int, error) {
	// I really REALLY don't like appending the SQL strings, but we can't yet parameterize table names with the
	// go sql driver, so this is how we have to do it. It's in a controlled enough environment where hopefully any
	// risk of SQL injection is mitigated.
	query := fmt.Sprintf("SELECT COUNT(1) AS Count FROM %s WHERE PlayerID = ?;", table)

	row := db.QueryRow(query, playerID)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, err
}

func getTotalCount(db *sql.DB, table string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(1) as Count FROM %s;", table)

	row := db.QueryRow(query)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, err
}
