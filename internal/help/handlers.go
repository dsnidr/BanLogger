package help

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/banlogger/pkg/config"
	"github.com/sniddunc/gcmd"
)

// CommandHandler is the command handler for the warn command
func CommandHandler(c gcmd.Context) error {
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "Available Commands",
		Color: config.EmbedHelpColour,
		Description: `
		!warn <profileURL> <reason>

		!kick <profileURL> <reason>

		!ban <profileURL> <duration> <reason>
		Duration examples: 1min, 1h, 1d, 1w, 1m, 1y, perm

		!lookup <profileURL>

		Valid profileURL formats:
		https://steamcommunity.com/id/VANITY_URL/
		https://steamcommunity.com/profiles/STEAMID64/
		*The trailing slash is not required.
		`,
	})

	return nil
}
