package live

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/gcmd"
)

// HelpHandler is the command handler for the help command
func (handlers *CommandHandlers) HelpHandler(c gcmd.Context) error {
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "Available Commands",
		Color: config.EmbedHelpColour,
		Description: `
		!help
		!warn <profileURL> <reason>
		!kick <profileURL> <reason>
		!ban <profileURL> <duration> <reason>
		Duration examples: 1min, 3h, 1d, 10w, 1m, 2y, perm
		!lookup <profileURL>
		!top
		!stats

		Valid profileURL formats:
		https://steamcommunity.com/id/VANITY_URL/
		https://steamcommunity.com/profiles/STEAMID64/
		*The trailing slash is not required.
		`,
	})

	return nil
}
