package git_cmd_svc

import (
	"fmt"
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"

	real_os "os"
)

func TestGitFork(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	os.On("Command", "git", "checkout", "-b", "new-branch").Return([]byte("Switched to a new branch 'new-branch'"), nil)
	os.On("UserHomeDir").Return("/home/testuser", nil)
	os.On("MkdirAll", "/home/testuser/.git_project/repo", real_os.ModePerm).Return(nil)
	os.On("WriteFile", "/home/testuser/.git_project/repo/new-branch.parent", []byte("parent-branch"), real_os.FileMode(0644)).Return(nil)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("new-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(true, nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForGetBranch(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte(""), real_os.ErrInvalid)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for getting current branch, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForGetRepo(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte(""), real_os.ErrInvalid)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for getting repo name, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForRead(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("", real_os.ErrInvalid)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for reading new branch name, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForReadEof(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("", fmt.Errorf("EOF"))

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()
	if err != nil {
		t.Errorf("Expected no error for EOF, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0 for EOF, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrDuplicationBranch(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("current-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("current-branch", nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for duplicate branch name, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkConfirmCancel(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("new-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(false, nil)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for cancelled confirmation, got nil")
	}
	if code != 0 {
		t.Errorf("Expected exit code 0 for cancellation, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForConfirm(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("new-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(false, real_os.ErrInvalid)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for cancelled confirmation, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1 for error, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForCheckout(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("new-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(true, nil)
	os.On("Command", "git", "checkout", "-b", "new-branch").Return([]byte(""), real_os.ErrInvalid)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for git checkout failure, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForSetParentUserHomeDir(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("new-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(true, nil)
	os.On("Command", "git", "checkout", "-b", "new-branch").Return([]byte("Switched to a new branch 'new-branch'"), nil)
	os.On("UserHomeDir").Return("/home/testuser", real_os.ErrInvalid)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for UserHomeDir failure, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForSetParentMkdirAll(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("new-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(true, nil)
	os.On("Command", "git", "checkout", "-b", "new-branch").Return([]byte("Switched to a new branch 'new-branch'"), nil)
	os.On("UserHomeDir").Return("/home/testuser", nil)
	os.On("MkdirAll", "/home/testuser/.git_project/repo", real_os.ModePerm).Return(real_os.ErrInvalid)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for MkdirAll failure, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}

func TestGitForkErrForSetParentWriteFile(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("parent-branch\n"), nil)
	os.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	cli_pkg.On("Read", "新しいブランチ名を入力してください: ").Return("new-branch", nil)
	cli_pkg.On("Confirm", "この操作を実行しますか？").Return(true, nil)
	os.On("Command", "git", "checkout", "-b", "new-branch").Return([]byte("Switched to a new branch 'new-branch'"), nil)
	os.On("UserHomeDir").Return("/home/testuser", nil)
	os.On("MkdirAll", "/home/testuser/.git_project/repo", real_os.ModePerm).Return(nil)
	os.On("WriteFile", "/home/testuser/.git_project/repo/new-branch.parent", []byte("parent-branch"), real_os.FileMode(0644)).Return(real_os.ErrInvalid)

	s := NewGitCmdService(os, cli_pkg)
	code, err := s.GitFork()

	if err == nil {
		t.Errorf("Expected error for WriteFile failure, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}

	os.AssertExpectations(t)
	cli_pkg.AssertExpectations(t)
}
