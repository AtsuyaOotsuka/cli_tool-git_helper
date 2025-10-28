package svc

import (
	"testing"

	"github.com/atylab-libs/go_cli_tool-libs/pkg/utils_pkg"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"

	"github.com/stretchr/testify/assert"
)

func TestNewCheckService(t *testing.T) {
	os := utils_pkg.NewOS()
	checkService := NewCheckService(os)

	assert.NotNil(t, checkService)
	assert.Equal(t, os, checkService.os)
}

func TestCanStart(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "--version").Return([]byte("git version 2.30.0"), nil)
	os.On("Command", "git", "rev-parse", "--is-inside-work-tree").Return([]byte("true"), nil)
	checkService := NewCheckService(os)

	err := checkService.CanStart()
	assert.Nil(t, err)
	os.AssertExpectations(t)
}

func TestCanStart_GitNotInstalled(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "--version").Return([]byte(""), assert.AnError)
	checkService := NewCheckService(os)

	err := checkService.CanStart()
	assert.NotNil(t, err)
	assert.Equal(t, "🚫 Gitがインストールされていないようです。", err.Error())
	os.AssertExpectations(t)
}

func TestCanStart_NotInsideGitRepo(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "--version").Return([]byte("git version 2.30.0"), nil)
	os.On("Command", "git", "rev-parse", "--is-inside-work-tree").Return([]byte("false"), nil)
	checkService := NewCheckService(os)

	err := checkService.CanStart()
	assert.NotNil(t, err)
	assert.Equal(t, "🚫 Gitリポジトリ内で実行してください。", err.Error())
	os.AssertExpectations(t)
}

func TestCanStart_RevParseError(t *testing.T) {
	os := new(utils_pkg_mock.OsMock)
	os.On("Command", "git", "--version").Return([]byte("git version 2.30.0"), nil)
	os.On("Command", "git", "rev-parse", "--is-inside-work-tree").Return([]byte(""), assert.AnError)
	checkService := NewCheckService(os)

	err := checkService.CanStart()
	assert.NotNil(t, err)
	assert.Equal(t, "🚫 Gitリポジトリ内で実行してください。", err.Error())
	os.AssertExpectations(t)
}
