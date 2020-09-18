package mock

import (
	"strings"
	"time"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/gcmd"
)

// BanHandler handles a mock ban command
func (handlers *CommandHandlers) BanHandler(c gcmd.Context) error {
	duration := c.Get("duration").(string)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	ban := banlogger.Ban{
		PlayerID:  steamID,
		Duration:  duration,
		Reason:    reason,
		Staff:     strings.Repeat("1", 18),
		Timestamp: time.Now().Unix(),
	}

	handlers.BanService.CreateBan(ban)

	return nil
}
