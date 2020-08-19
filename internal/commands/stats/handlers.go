package stats

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/banlogger/internal/stats"
	"github.com/sniddunc/banlogger/pkg/config"
	"github.com/sniddunc/banlogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// CommandHandler is the command handler for the warn command
func CommandHandler(c gcmd.Context) error {
	db := c.Get("db").(*sql.DB)
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	totalWarnings, err := stats.GetTotalWarnCount(db)
	if err != nil {
		logging.Info("stats/handlers.go", fmt.Sprintf("GetTotalWarnCount returned an error: %v", err))
		totalWarnings = -1
	}

	totalKicks, err := stats.GetTotalKickCount(db)
	if err != nil {
		logging.Info("stats/handlers.go", fmt.Sprintf("GetTotalKickCount returned an error: %v", err))
		totalKicks = -1
	}

	totalBans, err := stats.GetTotalBanCount(db)
	if err != nil {
		logging.Info("stats/handlers.go", fmt.Sprintf("GetTotalBanCount returned an error: %v", err))
		totalBans = -1
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "BanLogger Stats",
		Color: config.EmbedHelpColour,
		Description: fmt.Sprintf(`
		Total Warnings: %d
		Total Kicks: %d
		Total Bans: %d
		`, totalWarnings, totalKicks, totalBans),
	})

	return nil
}
