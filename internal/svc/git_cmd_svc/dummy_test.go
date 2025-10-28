package git_cmd_svc

import (
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestGitCommitRedo(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommitRedo()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}
}
