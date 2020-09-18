package mock

import "github.com/sniddunc/gcmd"

// TopHandler handles a mock top command
func (handlers *CommandHandlers) TopHandler(c gcmd.Context) error {
	return nil
}
