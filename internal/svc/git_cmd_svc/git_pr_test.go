package git_cmd_svc

import (
	"os"
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"
)

func TestGitPr(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	os_mock.On("UserHomeDir").Return("/home/testuser", nil)
	file_info, _ := os.Stat(".")
	os_mock.On("Stat", "/home/testuser/.git_project/repo/feature-123.parent").Return(file_info, nil)
	os_mock.On("IsNotExist", nil).Return(false)
	os_mock.On("ReadFile", "/home/testuser/.git_project/repo/feature-123.parent").Return([]byte("main\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	os_mock.On("CommandRun", "open", "https://github.com/user/repo/compare/main...feature-123").Return(nil)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if code != 0 {
		t.Errorf("Expected exit code 0, got %d", code)
	}
}

func TestGitPrErrForGetCurrentBranchName(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), os.ErrNotExist)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitPrErrForGetRepoName(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), os.ErrNotExist)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitPrErrForGetParent(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	os_mock.On("UserHomeDir").Return("/home/testuser", os.ErrNotExist)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitPrErrForGetRemoteUrl(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil).Once()
	os_mock.On("UserHomeDir").Return("/home/testuser", nil)
	file_info, _ := os.Stat(".")
	os_mock.On("Stat", "/home/testuser/.git_project/repo/feature-123.parent").Return(file_info, nil)
	os_mock.On("IsNotExist", nil).Return(false)
	os_mock.On("ReadFile", "/home/testuser/.git_project/repo/feature-123.parent").Return([]byte("main\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), os.ErrNotExist).Once()

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitPrErrForOpenBrowser(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil).Once()
	os_mock.On("UserHomeDir").Return("/home/testuser", nil)
	file_info, _ := os.Stat(".")
	os_mock.On("Stat", "/home/testuser/.git_project/repo/feature-123.parent").Return(file_info, nil)
	os_mock.On("IsNotExist", nil).Return(false)
	os_mock.On("ReadFile", "/home/testuser/.git_project/repo/feature-123.parent").Return([]byte("main\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil).Once()
	os_mock.On("CommandRun", "open", "https://github.com/user/repo/compare/main...feature-123").Return(os.ErrNotExist)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitPrErrForNotFoundParent(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil).Once()
	os_mock.On("UserHomeDir").Return("/home/testuser", nil)
	file_info, _ := os.Stat(".")
	os_mock.On("Stat", "/home/testuser/.git_project/repo/feature-123.parent").Return(file_info, os.ErrNotExist)
	os_mock.On("IsNotExist", os.ErrNotExist).Return(true)
	cli_pkg := new(cli_pkg_mock.CliPkgMock)
	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitErrForReadFile(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	os_mock.On("UserHomeDir").Return("/home/testuser", nil)
	file_info, _ := os.Stat(".")
	os_mock.On("Stat", "/home/testuser/.git_project/repo/feature-123.parent").Return(file_info, nil)
	os_mock.On("IsNotExist", nil).Return(false)
	os_mock.On("ReadFile", "/home/testuser/.git_project/repo/feature-123.parent").Return([]byte("main\n"), os.ErrNotExist)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitPrErrForEmptyParent(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil)
	os_mock.On("UserHomeDir").Return("/home/testuser", nil)
	file_info, _ := os.Stat(".")
	os_mock.On("Stat", "/home/testuser/.git_project/repo/feature-123.parent").Return(file_info, nil)
	os_mock.On("IsNotExist", nil).Return(false)
	os_mock.On("ReadFile", "/home/testuser/.git_project/repo/feature-123.parent").Return([]byte("\n"), nil)

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}

func TestGitPrErrForGitUserFail(t *testing.T) {
	os_mock := new(utils_pkg_mock.OsMock)
	os_mock.On("Command", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("feature-123\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("https://github.com/user/repo.git\n"), nil).Once()
	os_mock.On("UserHomeDir").Return("/home/testuser", nil)
	file_info, _ := os.Stat(".")
	os_mock.On("Stat", "/home/testuser/.git_project/repo/feature-123.parent").Return(file_info, nil)
	os_mock.On("IsNotExist", nil).Return(false)
	os_mock.On("ReadFile", "/home/testuser/.git_project/repo/feature-123.parent").Return([]byte("main\n"), nil)
	os_mock.On("Command", "git", "config", "--get", "remote.origin.url").Return([]byte("invalid-url\n"), nil).Once()

	cli_pkg := new(cli_pkg_mock.CliPkgMock)

	s := NewGitCmdService(os_mock, cli_pkg)
	code, err := s.GitPr()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if code != 1 {
		t.Errorf("Expected exit code 1, got %d", code)
	}
}
