package top

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/banlogger/internal/stats"
	"github.com/sniddunc/banlogger/internal/steam"
	"github.com/sniddunc/banlogger/pkg/config"
	"github.com/sniddunc/banlogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// CommandHandler is the command handler for the warn command
func CommandHandler(c gcmd.Context) error {
	db := c.Get("db").(*sql.DB)
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	record, err := stats.GetPlayerWithMostInfractions(db)
	if err != nil {
		logging.Info("top/handlers.go", fmt.Sprintf("GetPlayerWithMostInfractions returned an error: %v", err))
		return fmt.Errorf("Could not get top player")
	}

	summary, err := steam.GetUserSummary(record.PlayerID)
	if err != nil {
		logging.Info("top/handlers.go", fmt.Sprintf("GetUserSummary returned an error: %v", err))
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
		`, summary.ProfileName, record.PlayerID, record.TotalInfractions, record.WarnCount, record.KickCount, record.BanCount, summary.ProfileURL),
		Image: &discordgo.MessageEmbedImage{
			URL:    summary.AvatarURL,
			Width:  100,
			Height: 100,
		},
	})

	return nil
}
