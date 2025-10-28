package git_cmd_svc

import (
	"fmt"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"

	"testing"
)

func TestGitCom(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("test-branch\n"), nil)

	cliMock.On("Read", "ディスクリプションを入力してください: ").Return("This is a test description", nil)

	osMock.On("Command", "git", "config", "branch.test-branch.description", "This is a test description").Return([]byte(""), nil)

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCom()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCom_EOF(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("test-branch\n"), nil)

	cliMock.On("Read", "ディスクリプションを入力してください: ").Return("", fmt.Errorf("EOF"))

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCom()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCom_GetBranchError(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte(""), fmt.Errorf("error"))

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCom()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCom_ConfigError(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("test-branch\n"), nil)

	cliMock.On("Read", "ディスクリプションを入力してください: ").Return("This is a test description", nil)

	osMock.On("Command", "git", "config", "branch.test-branch.description", "This is a test description").Return([]byte(""), fmt.Errorf("error"))

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCom()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}

func TestGitCom_ReadError(t *testing.T) {
	osMock := new(utils_pkg_mock.OsMock)
	cliMock := new(cli_pkg_mock.CliPkgMock)

	osMock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("test-branch\n"), nil)

	cliMock.On("Read", "ディスクリプションを入力してください: ").Return("", fmt.Errorf("some error"))

	gitCmdSvc := &GitCmdSvc{
		os:      osMock,
		cli_pkg: cliMock,
	}

	exitCode, err := gitCmdSvc.GitCom()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}

	osMock.AssertExpectations(t)
	cliMock.AssertExpectations(t)
}
