package mock

import (
	"log"

	"github.com/sniddunc/gcmd"
)

// BanListHandler is the command handler for the banlist command
func (handlers *CommandHandlers) BanListHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.BanListHandler"

	// s := c.Get("session").(*discordgo.Session)
	// m := c.Get("message").(*discordgo.MessageCreate)

	// s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
	// 	Title:       "Banned " + ban.PlayerID + " for " + ban.Reason,
	// 	Description: "Duration: " + ban.Duration + "\nBanned by: " + m.Author.Username + "\n" + helpers.GetInfractionString(handlers.StatService, steamID),
	// 	Color:       config.EmbedBanColour,
	// 	Footer: &discordgo.MessageEmbedFooter{
	// 		Text: ban.Staff,
	// 	},
	// })

	currentBans, err := handlers.BanService.GetCurrentBans()
	if err != nil {
		log.Println(err)
	}

	log.Println(currentBans)

	return nil
}
