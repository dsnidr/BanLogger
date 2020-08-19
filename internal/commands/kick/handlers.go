package kick

import (
	"database/sql"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/banlogger/pkg/config"
	"github.com/sniddunc/banlogger/pkg/helpers"
	"github.com/sniddunc/gcmd"
)

// CommandHandler is the command handler for the kick command
func CommandHandler(c gcmd.Context) error {
	db := c.Get("db").(*sql.DB)
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	kick := Kick{
		PlayerID:  steamID,
		Reason:    reason,
		Staff:     m.Author.ID,
		Timestamp: time.Now().Unix(),
	}

	err := Insert(db, kick)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Kicked " + kick.PlayerID + " for " + kick.Reason,
		Description: "Kicked by: " + m.Author.Username + "\n" + helpers.GetInfractionString(db, steamID),
		Color:       config.EmbedKickColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: kick.Staff,
		},
	})

	return nil
}
