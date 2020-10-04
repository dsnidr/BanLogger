package live

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/helpers"
	"github.com/sniddunc/gcmd"
)

// BanListHandler is the command handler for the banlist command
func (handlers *CommandHandlers) BanListHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.BanListHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	currentBans, err := handlers.BanService.GetCurrentBans()
	if err != nil {
		log.Println(err)
	}

	banlist := ""

	alreadyShown := []string{}

	s.ChannelMessageSend(m.ChannelID, "This might take a minute...")

	for _, ban := range currentBans {
		if helpers.ContainsString(alreadyShown, ban.PlayerID) {
			continue
		}

		// Check if this player's summary is already cached
		found, cached := handlers.PlayerSummaryCache.Get(ban.PlayerID)
		if !cached {
			// If the player's summary isn't cached, resolve it
			summary, err := handlers.SteamService.GetUserSummary(ban.PlayerID)
			if err != nil {
				banlist += "Couldn't resolve player name\n"
				continue
			}

			// Cache them
			handlers.PlayerSummaryCache.Set(ban.PlayerID, summary, cache.DefaultExpiration)

			found = summary
		}

		summary := found.(banlogger.SteamPlayerSummary)

		banlist += fmt.Sprintf("%s\n%s\n\n", ban.PlayerID, summary.ProfileName)

		alreadyShown = append(alreadyShown, ban.PlayerID)
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Currently banned players",
		Color:       3434475,
		Description: banlist,
	})

	return nil
}
