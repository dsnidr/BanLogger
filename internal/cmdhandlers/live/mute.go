package live

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/helpers"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/BanLogger/pkg/strdur"
	"github.com/sniddunc/gcmd"
)

// MuteHandler is the command handler for the mute command
func (handlers *CommandHandlers) MuteHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.MuteHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	duration := c.Get("duration").(string)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	timestamp := time.Now().Unix()

	var unmutedAt int64
	if duration != "perm" {
		unmutedAt = timestamp + strdur.GetDuration(duration)
	} else {
		unmutedAt = 0
	}

	mute := banlogger.Mute{
		PlayerID:  steamID,
		Duration:  duration,
		Reason:    reason,
		Staff:     m.Author.ID,
		UnmutedAt: unmutedAt,
		Timestamp: timestamp,
	}

	err := handlers.MuteService.CreateMute(mute)
	if err != nil {
		logging.Error(tag, fmt.Sprintf("CreateMute returned an error: %v", err))

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Muted " + mute.PlayerID + " for " + mute.Reason,
		Description: "Duration: " + mute.Duration + "\nBanned by: " + m.Author.Username + "\n" + helpers.GetInfractionString(handlers.StatService, steamID),
		Color:       config.EmbedWarningColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: mute.Staff,
		},
	})

	return nil
}
