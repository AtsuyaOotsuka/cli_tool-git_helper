package git_cmd_svc

import (
	"fmt"
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestGitFetch(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	cli_pkg.On("Read", "フェッチするブランチ名を入力してください: ").Return("feature-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(true, nil)
	os.On("Command", "git", "fetch", "origin", "feature-branch").Return([]byte("Fetched"), nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFetch()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitFetchReadErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	cli_pkg.On("Read", "フェッチするブランチ名を入力してください: ").Return("", fmt.Errorf("read error"))
	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFetch()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitFetchEOF(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	cli_pkg.On("Read", "フェッチするブランチ名を入力してください: ").Return("", fmt.Errorf("EOF"))
	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFetch()

	if err != nil {
		t.Errorf("Expected no error for EOF, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0 for EOF, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitFetchEmptyBranch(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	cli_pkg.On("Read", "フェッチするブランチ名を入力してください: ").Return("", nil)
	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFetch()

	if err == nil {
		t.Errorf("Expected error for empty branch name, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitFetchCancel(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	cli_pkg.On("Read", "フェッチするブランチ名を入力してください: ").Return("feature-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(false, nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFetch()

	if err == nil {
		t.Errorf("Expected cancellation error, got nil")
	}
	if code != 0 {
		t.Errorf("Expected exit code 0 for cancellation, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitFetchConfirmErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	cli_pkg.On("Read", "フェッチするブランチ名を入力してください: ").Return("feature-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(false, fmt.Errorf("confirm error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFetch()

	if err == nil {
		t.Errorf("Expected error from Confirm, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitFetchGitErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	cli_pkg.On("Read", "フェッチするブランチ名を入力してください: ").Return("feature-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(true, nil)
	os.On("Command", "git", "fetch", "origin", "feature-branch").Return([]byte(""), fmt.Errorf("git fetch error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFetch()

	if err == nil {
		t.Errorf("Expected error from git fetch, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}
