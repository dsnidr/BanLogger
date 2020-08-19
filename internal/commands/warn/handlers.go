package warn

import (
	"database/sql"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/banlogger/pkg/config"
	"github.com/sniddunc/banlogger/pkg/helpers"
	"github.com/sniddunc/gcmd"
)

// CommandHandler is the command handler for the warn command
func CommandHandler(c gcmd.Context) error {
	db := c.Get("db").(*sql.DB)
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	warning := Warning{
		PlayerID:  steamID,
		Reason:    reason,
		Staff:     m.Author.ID,
		Timestamp: time.Now().Unix(),
	}

	err := Insert(db, warning)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Warned " + warning.PlayerID + " for " + warning.Reason,
		Description: "Warned by: " + m.Author.Username + "\n" + helpers.GetInfractionString(db, steamID),
		Color:       config.EmbedWarningColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: warning.Staff,
		},
	})

	return nil
}
