package git_cmd_svc

import (
	"fmt"
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestGitPull(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を pull しますか？").Return(true, nil)
	os.On("Command", "git", "pull", "origin", "feature-branch").Return([]byte("Pull successful"), nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPull()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPullGetBranchErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte(""), fmt.Errorf("branch error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPull()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPullCancel(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を pull しますか？").Return(false, nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPull()

	if err == nil {
		t.Errorf("Expected error for cancellation, got nil")
	}
	if code != 0 {
		t.Errorf("Expected exit code 0 for cancellation, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPullConfirmErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を pull しますか？").Return(false, fmt.Errorf("confirm error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPull()

	if err == nil {
		t.Errorf("Expected error from Confirm, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitPullGitErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-branch"), nil)
	cli_pkg.On("Confirm", "ブランチ feature-branch を pull しますか？").Return(true, nil)
	os.On("Command", "git", "pull", "origin", "feature-branch").Return([]byte(""), fmt.Errorf("git pull error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitPull()

	if err == nil {
		t.Errorf("Expected error from git pull, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}
