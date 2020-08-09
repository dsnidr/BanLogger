package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sniddunc/banlogger/internal/bot"
	"github.com/sniddunc/banlogger/pkg/logging"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment variables. Error: %v", err)
	}

	// Check existence of env variables
	botToken, ok := os.LookupEnv("DISCORD_BOT_TOKEN")
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

	// Setup database
	db := setupDatabase()

	// Setup bot functionality (commands, etc)
	bot.Setup(db)

	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("Could not create discord bot. Error: %v", err)
	}

	// Specify intents
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Setup event handlers
	dg.AddHandler(bot.MessageReceiveHandler)

	// Open socket to discord
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection to discord. Error: %v", err)
	}

	logging.Info("main.go", "Bot started")

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
		Timestamp INTEGER NOT NULL
	)`); err != nil {
		tx.Rollback()
		log.Fatalf("Could not create Ban table. Error: %v", err)
	}

	tx.Commit()

	return db
}
