package mock

import (
	"strings"
	"time"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/gcmd"
)

// KickHandler handles a mock help command
func (handlers *CommandHandlers) KickHandler(c gcmd.Context) error {
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	kick := banlogger.Kick{
		PlayerID:  steamID,
		Reason:    reason,
		Staff:     strings.Repeat("1", 18),
		Timestamp: time.Now().Unix(),
	}

	handlers.KickService.CreateKick(kick)

	return nil
}
