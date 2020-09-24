package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sniddunc/BanLogger/internal/bot"
	"github.com/sniddunc/BanLogger/internal/cmdhandlers/live"
	steamlive "github.com/sniddunc/BanLogger/internal/steam/live"
	"github.com/sniddunc/BanLogger/internal/storage/sqlite"
	"github.com/sniddunc/BanLogger/pkg/logging"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	// Check existence of env variables
	_, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
	if !ok {
		log.Fatalf("Environment variables DISCORD_BOT_TOKEN is required, but is not currently set. Exiting...")
	}

	_, ok = os.LookupEnv("STEAM_API_KEY")
	if !ok {
		log.Fatalf("Environment variables STEAM_API_KEY is required, but is not currently set. Exiting...")
	}

	_, ok = os.LookupEnv("CHANNEL_ID")
	if !ok {
		log.Fatalf("Environment variables CHANNEL_ID is required, but is not currently set. Exiting...")
	}

	// Database setup
	db := setupDatabase()
	defer db.Close()

	// Services setup
	steamService := &steamlive.SteamService{}
	warningService := &sqlite.WarningService{DB: db}
	kickService := &sqlite.KickService{DB: db}
	banService := &sqlite.BanService{DB: db}
	statService := &sqlite.StatService{
		DB:             db,
		WarningService: warningService,
		KickService:    kickService,
		BanService:     banService,
	}

	// Bot setup
	commandHandlers := &live.CommandHandlers{
		SteamService:   steamService,
		WarningService: warningService,
		KickService:    kickService,
		BanService:     banService,
		StatService:    statService,
	}

	bot := bot.Bot{
		SteamService:    steamService,
		CommandHandlers: commandHandlers,
	}

	dg, err := bot.Setup()
	if err != nil {
		log.Fatalf("Could not setup bot. Error: %v", err)
	}

	// Open bot socket
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening a connection to discord. Error: %v", err)
	}

	logging.Info("main.go", "Bot connected")

	// Wait for exit signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up and close session
	dg.Close()
}

func setupDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./banlogger.db")
	if err != nil {
		log.Fatalf("Could not establish database connection. Error: %v", err)
	}

	// Create tables
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("Could not start transaction. Error: %v", err)
	}

	// Create warning table
	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS Warning(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		PlayerID VARCHAR(17) NOT NULL,
		Reason VARCHAR(128) NOT NULL,
		Staff VARCHAR(64) NOT NULL,
		Timestamp INTEGER NOT NULL
	)`); err != nil {
		tx.Rollback()
		log.Fatalf("Could not create Warning table. Error: %v", err)
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS Kick(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		PlayerID VARCHAR(17) NOT NULL,
		Reason VARCHAR(128) NOT NULL,
		Staff VARCHAR(64) NOT NULL,
		Timestamp INTEGER NOT NULL
	)`); err != nil {
		tx.Rollback()
		log.Fatalf("Could not create Kick table. Error: %v", err)
	}

	if _, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS Ban(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		PlayerID VARCHAR(17) NOT NULL,
		Duration VARCHAR(6) NOT NULL,
		Reason VARCHAR(128) NOT NULL,
		Staff VARCHAR(64) NOT NULL,
		UnbannedAt INTEGER NOT NULL DEFAULT 0,
		Timestamp INTEGER NOT NULL
	)`); err != nil {
		tx.Rollback()
		log.Fatalf("Could not create Ban table. Error: %v", err)
	}

	tx.Commit()

	return db
}
