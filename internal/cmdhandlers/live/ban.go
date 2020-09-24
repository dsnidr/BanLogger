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

// BanHandler is the command handler for the ban command
func (handlers *CommandHandlers) BanHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.BanHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	duration := c.Get("duration").(string)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	timestamp := time.Now().Unix()

	var unbannedAt int64
	if duration != "perm" {
		unbannedAt = timestamp + strdur.GetDuration(duration)
	} else {
		unbannedAt = 0
	}

	ban := banlogger.Ban{
		PlayerID:   steamID,
		Duration:   duration,
		Reason:     reason,
		Staff:      m.Author.ID,
		UnbannedAt: unbannedAt,
		Timestamp:  timestamp,
	}

	err := handlers.BanService.CreateBan(ban)
	if err != nil {
		logging.Error(tag, fmt.Sprintf("CreateBan returned an error: %v", err))

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Banned " + ban.PlayerID + " for " + ban.Reason,
		Description: "Duration: " + ban.Duration + "\nBanned by: " + m.Author.Username + "\n" + helpers.GetInfractionString(handlers.StatService, steamID),
		Color:       config.EmbedBanColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: ban.Staff,
		},
	})

	return nil
}
