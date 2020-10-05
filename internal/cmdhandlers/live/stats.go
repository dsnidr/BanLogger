package live

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// StatsHandler is the command handler for the warn command
func (handlers *CommandHandlers) StatsHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.StatsHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	totalMutes, err := handlers.StatService.GetTotalMuteCount()
	if err != nil {
		logging.Info(tag, fmt.Sprintf("GetTotalMuteCount returned an error: %v", err))
		totalMutes = -1
	}

	totalWarnings, err := handlers.StatService.GetTotalWarningCount()
	if err != nil {
		logging.Info(tag, fmt.Sprintf("GetTotalWarningCount returned an error: %v", err))
		totalWarnings = -1
	}

	totalKicks, err := handlers.StatService.GetTotalKickCount()
	if err != nil {
		logging.Info(tag, fmt.Sprintf("GetTotalKickCount returned an error: %v", err))
		totalKicks = -1
	}

	totalBans, err := handlers.StatService.GetTotalBanCount()
	if err != nil {
		logging.Info(tag, fmt.Sprintf("GetTotalBanCount returned an error: %v", err))
		totalBans = -1
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "BanLogger Stats",
		Color: config.EmbedHelpColour,
		Description: fmt.Sprintf(`
		Total Warnings: %d
		Total Mutes: %d
		Total Kicks: %d
		Total Bans: %d
		`, totalWarnings, totalMutes, totalKicks, totalBans),
	})

	return nil
}
