package live

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/helpers"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// KickHandler is the handler for the kick command
func (handlers *CommandHandlers) KickHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.KickHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	kick := banlogger.Kick{
		PlayerID:  steamID,
		Reason:    reason,
		Staff:     m.Author.ID,
		Timestamp: time.Now().Unix(),
	}

	err := handlers.KickService.CreateKick(kick)
	if err != nil {
		logging.Error(tag, fmt.Sprintf("CreateKick returned an error: %v", err))

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Kicked " + kick.PlayerID + " for " + kick.Reason,
		Description: "Kicked by: " + m.Author.Username + "\n" + helpers.GetInfractionString(handlers.StatService, steamID),
		Color:       config.EmbedKickColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: kick.Staff,
		},
	})

	return nil
}
