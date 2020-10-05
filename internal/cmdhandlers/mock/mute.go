package mock

import (
	"strings"
	"time"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	"github.com/sniddunc/gcmd"
)

// MuteHandler handles a mock mute command
func (handlers *CommandHandlers) MuteHandler(c gcmd.Context) error {
	duration := c.Get("duration").(string)
	reason := c.Get("reason").(string)
	steamID := c.Get("steamID").(string)

	mute := banlogger.Mute{
		PlayerID:  steamID,
		Duration:  duration,
		Reason:    reason,
		Staff:     strings.Repeat("1", 18),
		Timestamp: time.Now().Unix(),
	}

	handlers.MuteService.CreateMute(mute)

	return nil
}
