package live

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// TopHandler is the command handler for the warn command
func (handlers *CommandHandlers) TopHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.TopHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	record, err := handlers.StatService.GetTopOffender()
	if err != nil {
		logging.Info(tag, fmt.Sprintf("GetPlayerWithMostInfractions returned an error: %v", err))
		return fmt.Errorf("Could not get top player")
	}

	summary, err := handlers.SteamService.GetUserSummary(record.PlayerID)
	if err != nil {
		logging.Info(tag, fmt.Sprintf("GetUserSummary returned an error: %v", err))
		return fmt.Errorf("Could not get top player")
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "Worst Offender",
		Color: config.EmbedHelpColour,
		Description: fmt.Sprintf(`
		Name: %s
		SteamID: %s
		Number of infractions: %d
		Warn count: %d
		Kick count: %d
		Ban count: %d
		Profile URL: %s
		`, summary.ProfileName, record.PlayerID, record.Total, record.WarningCount, record.KickCount, record.BanCount, summary.ProfileURL),
		Image: &discordgo.MessageEmbedImage{
			URL:    summary.AvatarURL,
			Width:  100,
			Height: 100,
		},
	})

	return nil
}
