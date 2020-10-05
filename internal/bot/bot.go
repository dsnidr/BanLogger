package bot

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/BanLogger/pkg/config"
	"github.com/sniddunc/BanLogger/pkg/logging"
	"github.com/sniddunc/gcmd"
)

// Bot holds all our services to be used in command handlers
type Bot struct {
	CmdBase         gcmd.Base
	SteamService    banlogger.SteamService
	CommandHandlers banlogger.CommandHandlers
}

// Setup initializes everything the bot needs to run
func (bot *Bot) Setup() (*discordgo.Session, error) {
	// gcmd setup
	cmdBase := gcmd.New("!")
	cmdBase.UnknownCommandMessage = "Unknown command. Do !help for a list of commands."

	// Register help command
	helpCommand := gcmd.Command{
		Name:    "help",
		Usage:   "!help",
		Handler: bot.CommandHandlers.HelpHandler,
	}
	cmdBase.Register(helpCommand)

	// Register warn command
	warnCommand := gcmd.Command{
		Name:    "warn",
		Usage:   "!warn <profileURL> <reason>",
		Handler: bot.CommandHandlers.WarnHandler,
	}
	warnCommand.Use(bot.ParseWarnArgs)
	cmdBase.Register(warnCommand)

	// Register kick command
	kickCommand := gcmd.Command{
		Name:    "kick",
		Usage:   "!kick <profileURL> <reason>",
		Handler: bot.CommandHandlers.KickHandler,
	}
	kickCommand.Use(bot.ParseKickArgs)
	cmdBase.Register(kickCommand)

	// Register ban command
	banCommand := gcmd.Command{
		Name:    "ban",
		Usage:   "!ban <profileURL> <reason>",
		Handler: bot.CommandHandlers.BanHandler,
	}
	banCommand.Use(bot.ParseBanArgs)
	cmdBase.Register(banCommand)

	// Register mute command
	muteCommand := gcmd.Command{
		Name:    "mute",
		Usage:   "!mute <profileURL> <reason>",
		Handler: bot.CommandHandlers.MuteHandler,
	}
	muteCommand.Use(bot.ParseMuteArgs)
	cmdBase.Register(muteCommand)

	// Register banlist command
	banlistCommand := gcmd.Command{
		Name:    "banlist",
		Usage:   "!banlist",
		Handler: bot.CommandHandlers.BanListHandler,
	}
	cmdBase.Register(banlistCommand)

	// Register lookup command
	lookupCommand := gcmd.Command{
		Name:    "lookup",
		Usage:   "!lookup <profileURL>",
		Handler: bot.CommandHandlers.LookupHandler,
	}
	lookupCommand.Use(bot.ParseLookupArgs)
	cmdBase.Register(lookupCommand)

	// Register stats command
	statsCommand := gcmd.Command{
		Name:    "stats",
		Usage:   "!stats",
		Handler: bot.CommandHandlers.StatsHandler,
	}
	cmdBase.Register(statsCommand)

	// Register top command
	topCommand := gcmd.Command{
		Name:    "top",
		Usage:   "!top",
		Handler: bot.CommandHandlers.TopHandler,
	}
	cmdBase.Register(topCommand)

	bot.CmdBase = cmdBase

	// Discordgo setup
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		return nil, err
	}

	// Provide intents
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Setup event handlers
	dg.AddHandler(bot.MessageReceivedHandler)

	return dg, nil
}

// MessageReceivedHandler is the handler for when the bot receives a message
func (bot *Bot) MessageReceivedHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages which did not originate from the provided channel
	if m.ChannelID != os.Getenv("CHANNEL_ID") {
		return
	}

	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Log received message
	logging.Info("bot/bot.go", fmt.Sprintf("Message received from %s (%s):\n\tMessage: '%s'", m.Author.Username, m.Author.ID, m.Content))

	// Pass along the session and message pointers to the command handler
	extraStore := map[string]interface{}{
		"session": s,
		"message": m,
	}

	_, err := bot.CmdBase.HandleCommand(m.Content, extraStore)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: m.Author.Mention(),
			Color:       config.EmbedErrorColour,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
