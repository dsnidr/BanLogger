package bot

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/banlogger/pkg/config"
)

// MessageReceiveHandler is the handler for when the bot receives a message
func MessageReceiveHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages which did not originate from the provided channel
	if m.ChannelID != os.Getenv("CHANNEL_ID") {
		return
	}

	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Pass along the session and message pointers to the command handler
	extraStore := map[string]interface{}{
		"session": s,
		"message": m,
	}

	_, err := cmdBase.HandleCommand(m.Content, extraStore)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
