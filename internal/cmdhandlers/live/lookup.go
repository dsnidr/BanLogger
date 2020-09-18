package live

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/pkg/helpers"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// LookupHandler is the handler for the lookup command
func (handlers *CommandHandlers) LookupHandler(c gcmd.Context) error {
	const tag = "cmdhandlers.live.LookupHandler"

	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	steamID := c.Get("steamID").(string)

	record, err := handlers.StatService.GetRecord(steamID)
	if err != nil {
		return err
	}

	summary, err := handlers.SteamService.GetUserSummary(record.PlayerID)
	if err != nil {
		logging.Info(tag, fmt.Sprintf("GetUserSummary returned an error: %v", err))
		return fmt.Errorf("Could not fetch record")
	}

	// Build record display
	warningDisplay := ""
	kickDisplay := ""
	banDisplay := ""

	if len(record.Warnings) == 0 {
		warningDisplay = "```No warnings found```"
	} else {
		for _, warning := range record.Warnings {
			timestamp := helpers.GetTimestamp(warning.Timestamp)
			warningDisplay += fmt.Sprintf("```ID: %-4d        %s\n> %s```", warning.ID, timestamp, warning.Reason)
		}
	}

	if len(record.Kicks) == 0 {
		kickDisplay = "```No kicks found```"
	} else {
		for _, kick := range record.Kicks {
			timestamp := helpers.GetTimestamp(kick.Timestamp)
			kickDisplay += fmt.Sprintf("```ID: %-4d        %s\n> %s```", kick.ID, timestamp, kick.Reason)
		}
	}

	if len(record.Bans) == 0 {
		banDisplay = "```No bans found```"
	} else {
		for _, ban := range record.Bans {
			timestamp := helpers.GetTimestamp(ban.Timestamp)
			banDisplay += fmt.Sprintf("```ID: %-4d %-6s %s\n> %s```", ban.ID, ban.Duration, timestamp, ban.Reason)
		}
	}

	recordDisplay := fmt.Sprintf("**Warnings:**\n%s\n**Kicks:**:\n%s\n**Bans:**\n%s\n", warningDisplay, kickDisplay, banDisplay)
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Showing " + summary.ProfileName + "'s record",
		Color:       3434475,
		Description: recordDisplay + "\nProfile: " + summary.ProfileURL,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "SteamID: " + record.PlayerID,
		},
	})

	return nil
}
