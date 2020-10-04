package live

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/helpers"
	"github.com/sniddunc/gcmd"
)

type banSummary struct {
	PlayerID    string
	Duration    string
	Name        string
	MinutesLeft int
}

const maxFieldsPerEmbed = 25

// BanListHandler is the command handler for the banlist command
func (handlers *CommandHandlers) BanListHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.BanListHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	excludePermBans := false
	if len(c.Args) > 0 && strings.ToLower(c.Args[0]) == "-ep" {
		excludePermBans = true
	}

	currentBans, err := handlers.BanService.GetCurrentBans()
	if err != nil {
		log.Println(err)
	}

	alreadyShown := []string{}
	banSummaries := []banSummary{}

	s.ChannelMessageSend(m.ChannelID, "This might take a minute...")

	// Gather banned player summaries
	for _, ban := range currentBans {
		if excludePermBans && ban.Duration == "perm" {
			continue
		}

		if helpers.ContainsString(alreadyShown, ban.PlayerID) {
			continue
		}

		// Check if this player's summary is already cached
		found, cached := handlers.PlayerSummaryCache.Get(ban.PlayerID)
		if !cached {
			// If the player's summary isn't cached, resolve it
			summary, err := handlers.SteamService.GetUserSummary(ban.PlayerID)
			if err != nil {
				log.Printf("Couldn't retrieve player %s summary. Error: %v", ban.PlayerID, err)
				continue
			}

			// Cache them
			handlers.PlayerSummaryCache.Set(ban.PlayerID, summary, cache.DefaultExpiration)

			found = summary
		}

		summary := found.(banlogger.SteamPlayerSummary)

		// Determine minutes remaining in ban
		var minutesLeft int64
		if ban.Duration == "perm" {
			minutesLeft = -1
		} else {
			minutesLeft = (ban.UnbannedAt - time.Now().Unix()) / 60
		}

		fmt.Printf("%s | Now: %d, UnbannedAt: %d\n", summary.ProfileName, time.Now().Unix(), ban.UnbannedAt)

		banSummary := banSummary{
			PlayerID:    ban.PlayerID,
			Duration:    ban.Duration,
			Name:        summary.ProfileName,
			MinutesLeft: int(minutesLeft),
		}

		banSummaries = append(banSummaries, banSummary)
		alreadyShown = append(alreadyShown, ban.PlayerID)
	}

	// Build output fields
	fields := []*discordgo.MessageEmbedField{}
	titleShown := false

	embedTitle := "Currently banned players"
	embedDescription := "Legend:\nname, steamID\nduration, minutes left in ban\n\n-1 minutes left indicates a perm ban"

	for i := 0; i < len(banSummaries); i++ {
		summary := banSummaries[i]

		field := &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("```%-20s %-17s```", summary.Name, summary.PlayerID),
			Value: fmt.Sprintf("```%-20s %-20s```", summary.Duration, fmt.Sprintf("for %d mins", summary.MinutesLeft)),
		}

		fields = append(fields, field)

		if len(fields) >= 25 {
			// Send message now and clear fields to continue
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       embedTitle,
				Description: embedDescription,
				Color:       3434475,
				Fields:      fields,
			})

			fields = []*discordgo.MessageEmbedField{}

			// Only show the title and description once
			if !titleShown {
				embedTitle = ""
				embedDescription = ""
				titleShown = true
			}
		}
	}

	// Send last embed
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       embedTitle,
		Description: embedDescription,
		Color:       3434475,
		Fields:      fields,
	})

	return nil
}
