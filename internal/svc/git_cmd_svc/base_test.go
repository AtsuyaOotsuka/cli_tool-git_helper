package git_cmd_svc

import (
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestNewGitCmdService(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	gitCmdSvc := NewGitCmdService(osMock, cliMock)

	if gitCmdSvc == nil {
		t.Errorf("Expected GitCmdSvcInterface instance, got nil")
	}
}
