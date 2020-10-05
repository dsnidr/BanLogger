package live

import (
	"github.com/patrickmn/go-cache"
	"github.com/sniddunc/BanLogger/internal/banlogger"
)

// CommandHandlers is a struct which we attach live command handlers to in order to satisfy
// the requirements of cmdhandlers.CommandHandlers
type CommandHandlers struct {
	SteamService       banlogger.SteamService
	WarningService     banlogger.WarningService
	KickService        banlogger.KickService
	BanService         banlogger.BanService
	MuteService        banlogger.MuteService
	StatService        banlogger.StatService
	PlayerSummaryCache *cache.Cache
}
