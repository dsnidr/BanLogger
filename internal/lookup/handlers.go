package lookup

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/gcmd"
)

// CommandHandler is the command handler for the warn command
func CommandHandler(c gcmd.Context) error {
	db := c.Get("db").(*sql.DB)
	s := c.Get("session").(*discordgo.Session)
	m := c.Get("message").(*discordgo.MessageCreate)
	steamID := c.Get("steamID").(string)

	record, err := GetRecord(db, steamID)
	if err != nil {
		return err
	}

	// Build record display
	warningDisplay := ""
	kickDisplay := ""
	banDisplay := ""

	if len(record.Warnings) == 0 {
		warningDisplay = "```No warnings found```"
	} else {
		for _, warning := range record.Warnings {
			timestamp := getTimestamp(warning.Timestamp)
			warningDisplay += fmt.Sprintf("```ID: %-4d        %s\n> %s```", warning.ID, timestamp, warning.Reason)
		}
	}

	if len(record.Kicks) == 0 {
		kickDisplay = "```No kicks found```"
	} else {
		for _, kick := range record.Kicks {
			timestamp := getTimestamp(kick.Timestamp)
			kickDisplay += fmt.Sprintf("```ID: %-4d        %s\n> %s```", kick.ID, timestamp, kick.Reason)
		}
	}

	if len(record.Bans) == 0 {
		banDisplay = "```No bans found```"
	} else {
		for _, ban := range record.Bans {
			timestamp := getTimestamp(ban.Timestamp)
			banDisplay += fmt.Sprintf("```ID: %-4d %-6s %s\n> %s```", ban.ID, ban.Duration, timestamp, ban.Reason)
		}
	}

	recordDisplay := fmt.Sprintf("**Warnings:**\n%s\n**Kicks:**:\n%s\n**Bans:**\n%s\n", warningDisplay, kickDisplay, banDisplay)
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Showing " + record.PlayerID + "'s record",
		Color:       3434475,
		Description: recordDisplay,
	})

	return nil
}

func getTimestamp(unixtime int64) string {
	return time.Unix(unixtime, 0).Format("02/01/06 03:04:05 PM")
}
