package live

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/pkg/config"
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

	for _, ban := range currentBans {
		var currentName string

		playerSummary, err := handlers.SteamService.GetUserSummary(ban.PlayerID)
		if err != nil {
			currentName = "Could not resolve current name"
		}

		currentName = playerSummary.ProfileName

		banlist += currentName + "\n"
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Banlist",
		Description: banlist,
		Color:       config.EmbedLookupColour,
	})

	return nil
}
