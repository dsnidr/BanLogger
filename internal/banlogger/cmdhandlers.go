package banlogger

import "github.com/sniddunc/gcmd"

// CommandHandlers is an interface to represent gcmd command handler functions
type CommandHandlers interface {
	HelpHandler(gcmd.Context) error
	WarnHandler(gcmd.Context) error
	KickHandler(gcmd.Context) error
	BanHandler(gcmd.Context) error
	MuteHandler(gcmd.Context) error
	BanListHandler(gcmd.Context) error
	LookupHandler(gcmd.Context) error
	StatsHandler(gcmd.Context) error
	TopHandler(gcmd.Context) error
}
