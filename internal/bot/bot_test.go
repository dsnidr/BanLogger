package bot

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sniddunc/BanLogger/internal/banlogger"
	cmdmock "github.com/sniddunc/BanLogger/internal/cmdhandlers/mock"
	steammock "github.com/sniddunc/BanLogger/internal/steam/mock"
	"github.com/sniddunc/BanLogger/internal/storage/mock"
	"github.com/sniddunc/BanLogger/pkg/config"
)

var testBot Bot
var warningService banlogger.WarningService
var kickService banlogger.KickService
var banService banlogger.BanService
var steamService banlogger.SteamService

var testWarnCommands map[string]bool

func init() {
	warningService = &mock.WarningService{
		Warnings: []banlogger.Warning{},
	}
	kickService = &mock.KickService{
		Kicks: []banlogger.Kick{},
	}
	banService = &mock.BanService{
		Bans: []banlogger.Ban{},
	}
	steamService = &steammock.SteamService{}

	commandHandlers := &cmdmock.CommandHandlers{
		WarningService: warningService,
		KickService:    kickService,
		BanService:     banService,
		SteamService:   steamService,
	}

	testBot = Bot{
		SteamService:    steamService,
		CommandHandlers: commandHandlers,
	}

	testBot.Setup()

	testWarnCommands = map[string]bool{
		"!warn https://steamcommunity.com/id/sniddunc Team killing":                                              true,
		fmt.Sprintf("!warn https://steamcommunity.com/id/sniddunc %s", strings.Repeat("1", config.ReasonMinLen)): true, // min reason length
		fmt.Sprintf("!warn https://steamcommunity.com/id/sniddunc %s", strings.Repeat("1", config.ReasonMaxLen)): true, // max reason length
		"!warn https://steamcommunity.com/profiles/sniddunc Team killing":                                        false,
		fmt.Sprintf("!warn https://steamcommunity.com/profiles/%s Team killing", strings.Repeat("1", 1)):         false, // profile too short
		fmt.Sprintf("!warn https://steamcommunity.com/profiles/%s Team killing", strings.Repeat("1", 12)):        false, // profile mid length
		fmt.Sprintf("!warn https://steamcommunity.com/profiles/%s Team killing", strings.Repeat("1", 18)):        false, // profile too long
		fmt.Sprintf("!warn https://steamcommunity.com/profiles/%s Team killing", strings.Repeat("1", 17)):        true,  // profile proper length
	}
}

func TestWarningCreation(t *testing.T) {
	for cmd, shouldPass := range testWarnCommands {
		_, err := testBot.CmdBase.HandleCommand(cmd, nil)

		if err != nil && shouldPass {
			t.Errorf("Command '%s' should have passed but didn't. Error: %v\n", cmd, err)
		} else if err == nil && !shouldPass {
			t.Errorf("Command '%s' should have failed but didn't\n", cmd)
		}
	}
}
