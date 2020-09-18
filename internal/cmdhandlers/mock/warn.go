package mock

import (
	"strings"
	"time"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/gcmd"
)

// WarnHandler handles a mock warn command
func (handlers *CommandHandlers) WarnHandler(c gcmd.Context) error {
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	warn := banlogger.Warning{
		PlayerID:  steamID,
		Reason:    reason,
		Staff:     strings.Repeat("1", 18),
		Timestamp: time.Now().Unix(),
	}

	handlers.WarningService.CreateWarning(warn)

	return nil
}
