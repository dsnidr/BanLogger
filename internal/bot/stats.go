package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// StatsCommandHandler is the command handler for the warn command
func (bot *Bot) StatsCommandHandler(c gcmd.Context) error {
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	totalWarnings, err := bot.StatService.GetTotalWarningCount()
	if err != nil {
		logging.Info("bot/stats.go", fmt.Sprintf("GetTotalWarningCount returned an error: %v", err))
		totalWarnings = -1
	}

	totalKicks, err := bot.StatService.GetTotalKickCount()
	if err != nil {
		logging.Info("bot/stats.go", fmt.Sprintf("GetTotalKickCount returned an error: %v", err))
		totalKicks = -1
	}

	totalBans, err := bot.StatService.GetTotalBanCount()
	if err != nil {
		logging.Info("bot/stats.go", fmt.Sprintf("GetTotalBanCount returned an error: %v", err))
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
