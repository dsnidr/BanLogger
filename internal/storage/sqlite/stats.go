package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/logging"
)

// StatService is the struct which we attach functions to in order to
// satisfy the requirements of banlogger.StatService.
type StatService struct {
	DB             *sql.DB
	WarningService banlogger.WarningService
	KickService    banlogger.KickService
	BanService     banlogger.BanService
}

// GetWarningCount returns the number of warnings this user has received
func (s *StatService) GetWarningCount(playerID string) (int, error) {
	return getCountForPlayer(s.DB, "Warning", playerID)
}

// GetKickCount returns the number of kicks this user has received
func (s *StatService) GetKickCount(playerID string) (int, error) {
	return getCountForPlayer(s.DB, "Kick", playerID)
}

// GetBanCount returns the number of bans this user has received
func (s *StatService) GetBanCount(playerID string) (int, error) {
	return getCountForPlayer(s.DB, "Ban", playerID)
}

// GetTotalWarningCount returns the total number of warnings recorded
func (s *StatService) GetTotalWarningCount() (int, error) {
	return getTotalCount(s.DB, "Warning")
}

// GetTotalKickCount returns the total number of kicks recorded
func (s *StatService) GetTotalKickCount() (int, error) {
	return getTotalCount(s.DB, "Kick")
}

// GetTotalBanCount returns the total number of bans recorded
func (s *StatService) GetTotalBanCount() (int, error) {
	return getTotalCount(s.DB, "Ban")
}

// GetRecord retrieves a list of all offenses recorded for a player
func (s *StatService) GetRecord(playerID string) (banlogger.Record, error) {
	warnings, err := s.WarningService.GetWarningsByPlayerID(playerID)
	if err != nil {
		return banlogger.Record{}, err
	}

	kicks, err := s.KickService.GetKicksByPlayerID(playerID)
	if err != nil {
		return banlogger.Record{}, err
	}

	bans, err := s.BanService.GetBansByPlayerID(playerID)
	if err != nil {
		return banlogger.Record{}, err
	}

	record := banlogger.Record{
		PlayerID: playerID,
		Warnings: warnings,
		Kicks:    kicks,
		Bans:     bans,
	}

	logging.Info("sqlite/stats.go",
		fmt.Sprintf("Record retrieved for player %s.\n\tWarnings: %d | Kicks: %d | Bans: %d",
			playerID, len(record.Warnings), len(record.Kicks), len(record.Bans)))

	return record, nil
}

// GetTopOffender finds and retrieves the count of offenses for the top offending player
func (s *StatService) GetTopOffender() (banlogger.NumericRecord, error) {
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

	row := s.DB.QueryRow(query)

	var record banlogger.NumericRecord

	err := row.Scan(&record.PlayerID)
	if err != nil {
		return banlogger.NumericRecord{}, err
	}

	warningCount, err := s.GetWarningCount(record.PlayerID)
	if err != nil {
		return banlogger.NumericRecord{}, err
	}

	kickCount, err := s.GetKickCount(record.PlayerID)
	if err != nil {
		return banlogger.NumericRecord{}, err
	}

	banCount, err := s.GetBanCount(record.PlayerID)
	if err != nil {
		return banlogger.NumericRecord{}, err
	}

	record.Total = warningCount + kickCount + banCount
	record.WarningCount = warningCount
	record.KickCount = kickCount
	record.BanCount = banCount

	return record, err
}

// Helper functions to avoid code reuse
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
