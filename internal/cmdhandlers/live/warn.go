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

// WarnHandler is the handler for the warn command
func (handlers *CommandHandlers) WarnHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.WarnHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	warning := banlogger.Warning{
		PlayerID:  steamID,
		Reason:    reason,
		Staff:     m.Author.ID,
		Timestamp: time.Now().Unix(),
	}

	err := handlers.WarningService.CreateWarning(warning)
	if err != nil {
		logging.Error(tag, fmt.Sprintf("CreateWarning returned an error: %v", err))

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Warned " + warning.PlayerID + " for " + warning.Reason,
		Description: "Warned by: " + m.Author.Username + "\n" + helpers.GetInfractionString(handlers.StatService, steamID),
		Color:       config.EmbedWarningColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: warning.Staff,
		},
	})

	return nil
}
