package ban

import (
	"database/sql"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/banlogger/pkg/config"
	"github.com/sniddunc/banlogger/pkg/helpers"
	"github.com/sniddunc/gcmd"
)

// CommandHandler is the command handler for the ban command
func CommandHandler(c gcmd.Context) error {
	db := c.Get("db").(*sql.DB)
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	duration := c.Get("duration").(string)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	ban := Ban{
		PlayerID:  steamID,
		Duration:  duration,
		Reason:    reason,
		Staff:     m.Author.ID,
		Timestamp: time.Now().Unix(),
	}

	err := Insert(db, ban)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Banned " + ban.PlayerID + " for " + ban.Reason,
		Description: "Duration: " + ban.Duration + "\nBanned by: " + m.Author.Username + "\n" + helpers.GetInfractionString(db, steamID),
		Color:       config.EmbedBanColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: ban.Staff,
		},
	})

	return nil
}
