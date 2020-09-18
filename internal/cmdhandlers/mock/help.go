package mock

import "github.com/sniddunc/gcmd"

// HelpHandler handles a mock help command
func (handlers *CommandHandlers) HelpHandler(c gcmd.Context) error {
	return nil
}
