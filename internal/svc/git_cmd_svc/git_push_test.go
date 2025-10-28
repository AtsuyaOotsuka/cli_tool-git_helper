package git_cmd_svc

import (
	"fmt"
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestGitPush(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を push しますか？").Return(true, nil)
	os.On("Command", "git", "push", "origin", "feature-branch").Return([]byte("Push successful"), nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPush()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPushGetBranchErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte(""), fmt.Errorf("branch error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPush()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPushCancel(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を push しますか？").Return(false, nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPush()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPushConfirmErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を push しますか？").Return(false, fmt.Errorf("confirm error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPush()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPushCommandErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を push しますか？").Return(true, nil)
	os.On("Command", "git", "push", "origin", "feature-branch").Return([]byte(""), fmt.Errorf("push error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPush()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}
