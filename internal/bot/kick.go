package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/helpers"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/BanLogger/pkg/validation"
	"github.com/sniddunc/gcmd"
)

// KickCommandHandler is the handler for the kick command
func (bot *Bot) KickCommandHandler(c gcmd.Context) error {
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

	err := bot.KickService.CreateKick(kick)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Kicked " + kick.PlayerID + " for " + kick.Reason,
		Description: "Kicked by: " + m.Author.Username + "\n" + helpers.GetInfractionString(bot.StatService, steamID),
		Color:       config.EmbedKickColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: kick.Staff,
		},
	})

	return nil
}

// ParseKickArgs parses the received arguments for the kick command. It's main function is to
// detect errors, report these errors and if no errors occur, then it attaches the necessary data
// to the provided command context.
func (bot *Bot) ParseKickArgs(next gcmd.HandlerFunc) gcmd.HandlerFunc {
	return func(c gcmd.Context) error {
		args := c.Args

		// Check amount of arguments provided
		if len(args) < 2 {
			return fmt.Errorf("Invalid arguments")
		}

		profileURL := args[0]

		// Validate profileURL
		ok, isVanity := validation.ValidateProfileURL(profileURL)
		if !ok {
			return fmt.Errorf("Invalid profile URL provided")
		}

		reason := strings.Join(args[1:], " ")

		if len(reason) < config.ReasonMinLen || len(reason) > config.ReasonMaxLen {
			return fmt.Errorf("Reason must be between %d and %d in length", config.ReasonMinLen, config.ReasonMaxLen)
		}

		// Get the user's SteamID
		steamID, err := bot.SteamService.GetSteamID(isVanity, profileURL)
		if err != nil {
			return fmt.Errorf("Could not resolve SteamID")
		}

		// Map data to context store
		c.Set("reason", reason)
		c.Set("steamID", steamID)

		logging.Info("bot/kick.go", fmt.Sprintf("Kick command passed arguments check. steamID: %s | reason: %s", steamID, reason))

		return next(c)
	}
}
