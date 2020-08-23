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

// BanCommandHandler is the command handler for the ban command
func (bot *Bot) BanCommandHandler(c gcmd.Context) error {
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	duration := c.Get("duration").(string)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	ban := banlogger.Ban{
		PlayerID:  steamID,
		Duration:  duration,
		Reason:    reason,
		Staff:     m.Author.ID,
		Timestamp: time.Now().Unix(),
	}

	err := bot.BanService.CreateBan(ban)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		})
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Banned " + ban.PlayerID + " for " + ban.Reason,
		Description: "Duration: " + ban.Duration + "\nBanned by: " + m.Author.Username + "\n" + helpers.GetInfractionString(bot.StatService, steamID),
		Color:       config.EmbedBanColour,
		Footer: &discordgo.MessageEmbedFooter{
			Text: ban.Staff,
		},
	})

	return nil
}

// ParseBanArgs parses the received arguments for the ban command. It's main function is to
// detect errors, report these errors and if no errors occur, then it attaches the necessary data
// to the provided command context.
func (bot *Bot) ParseBanArgs(next gcmd.HandlerFunc) gcmd.HandlerFunc {
	return func(c gcmd.Context) error {
		args := c.Args

		// Check amount of arguments provided
		if len(args) < 3 {
			return fmt.Errorf("Invalid arguments")
		}

		profileURL := args[0]

		// Validate profileURL
		ok, isVanity := validation.ValidateProfileURL(profileURL)
		if !ok {
			return fmt.Errorf("Invalid profile URL provided")
		}

		duration := args[1]

		// Validate duration
		valid := validation.ValidateDuration(duration)
		if !valid {
			return fmt.Errorf("Invalid duration provided")
		}

		reason := strings.Join(args[2:], " ")

		// Validate reason
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
		c.Set("duration", duration)
		c.Set("steamID", steamID)

		logging.Info("bot/ban.go", fmt.Sprintf("Ban command passed arguments check. steamID: %s | duration: %s | reason: %s", steamID, duration, reason))

		return next(c)
	}
}
