package git_cmd_svc

import (
	"fmt"
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestGitCommitRedo(t *testing.T) {

	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "log", "-1", "--pretty=%s").Return([]byte("Initial commit"), nil)

	cliMock.On("Confirm", "このコミットを取り消しますか？").Return(true, nil)

	osMock.On("Command", "git", "reset", "--soft", "HEAD~1").Return([]byte(""), nil)

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCommitRedo()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCommitRedoGetCommitErr(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "log", "-1", "--pretty=%s").Return([]byte(""), fmt.Errorf("git log error"))

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCommitRedo()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCommitRedoCancel(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "log", "-1", "--pretty=%s").Return([]byte("Initial commit"), nil)

	cliMock.On("Confirm", "このコミットを取り消しますか？").Return(false, nil)

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCommitRedo()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCommitRedoConfirmErr(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "log", "-1", "--pretty=%s").Return([]byte("Initial commit"), nil)

	cliMock.On("Confirm", "このコミットを取り消しますか？").Return(false, fmt.Errorf("confirm error"))

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCommitRedo()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCommitRedoResetErr(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "log", "-1", "--pretty=%s").Return([]byte("Initial commit"), nil)

	cliMock.On("Confirm", "このコミットを取り消しますか？").Return(true, nil)

	osMock.On("Command", "git", "reset", "--soft", "HEAD~1").Return([]byte(""), fmt.Errorf("git reset error"))

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCommitRedo()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}
