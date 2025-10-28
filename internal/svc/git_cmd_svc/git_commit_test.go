package git_cmd_svc

import (
	"fmt"
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestGitCommit(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123-add-new-feature\n"), nil)
	cli_pkg.On("Read", "コミットメッセージを入力してください: ").Return("Add new feature", nil)
	os.On("Command", "git", "commit", "-m", "feature-123 Add new feature").Return([]byte("Committed"), nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommit()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}
}

func TestGitCommitNoPrefix(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("main\n"), nil)
	cli_pkg.On("Read", "コミットメッセージを入力してください: ").Return("Initial commit", nil)
	os.On("Command", "git", "commit", "-m", "Initial commit").Return([]byte("Committed"), nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommit()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}
}

func TestGitCommitEmptyMessage(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123-add-new-feature\n"), nil)
	cli_pkg.On("Read", "コミットメッセージを入力してください: ").Return("", nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommit()

	if err == nil {
		t.Errorf("Expected error for empty commit message, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}
}

func TestGitCommitReadError(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123-add-new-feature\n"), nil)
	cli_pkg.On("Read", "コミットメッセージを入力してください: ").Return("", fmt.Errorf("Some read error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommit()

	if err == nil {
		t.Errorf("Expected error from Read, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}
}

func TestGitCommitGetBranchError(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte(""), fmt.Errorf("Some git error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommit()

	if err == nil {
		t.Errorf("Expected error from getCurrentBranchName, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}
}

func TestGitCommitReadEof(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123-add-new-feature\n"), nil)
	cli_pkg.On("Read", "コミットメッセージを入力してください: ").Return("", fmt.Errorf("EOF"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommit()

	if err != nil {
		t.Errorf("Expected no error on EOF, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0 on EOF, got %d", code)
	}
}

func TestGetCommitCommitErr(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123-add-new-feature\n"), nil)
	cli_pkg.On("Read", "コミットメッセージを入力してください: ").Return("Add new feature", nil)
	os.On("Command", "git", "commit", "-m", "feature-123 Add new feature").Return([]byte(""), fmt.Errorf("Some git commit error"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitCommit()

	if err == nil {
		t.Errorf("Expected error from git commit, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}
}
