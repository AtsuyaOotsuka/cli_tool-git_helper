package svc

import (
	"github.com/atylab-libs/go_cli_tool-libs/pkg/utils_pkg"

	"fmt"
	"strings"
)

type CheckSvcInterface interface {
	CanStart() error
}

type CheckSvcStruct struct {
	os utils_pkg.OsInterface
}

func NewCheckService(os utils_pkg.OsInterface) *CheckSvcStruct {
	return &CheckSvcStruct{
		os: os,
	}
}

func isInstallGit(cs *CheckSvcStruct) error {
	_, err := cs.os.Command("git", "--version")
	if err != nil {
		return fmt.Errorf("🚫 Gitがインストールされていないようです。")
	}
	return nil
}

// Gitリポジトリ内か確認
func isInsideGitRepo(cs *CheckSvcStruct) error {
	out, err := cs.os.Command("git", "rev-parse", "--is-inside-work-tree")
	if err != nil {
		return fmt.Errorf("🚫 Gitリポジトリ内で実行してください。")
	}
	result := strings.TrimSpace(string(out))

	if result != "true" {
		return fmt.Errorf("🚫 Gitリポジトリ内で実行してください。")
	}
	return nil
}

// 起動可能か確認
func (cs *CheckSvcStruct) CanStart() error {
	if err := isInstallGit(cs); err != nil {
		return err
	}
	if err := isInsideGitRepo(cs); err != nil {
		return err
	}
	return nil
}
