package app

import (
	"flag"
	"fmt"

	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/config"
	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/svc"
	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/svc/git_cmd_svc"

	"github.com/atylab-libs/go_cli_tool-libs/pkg/cli_pkg"
	"github.com/atylab-libs/go_cli_tool-libs/pkg/utils_pkg"
)

type App struct {
	os           utils_pkg.OsInterface
	check        svc.CheckSvcInterface
	git_cmd_svc  git_cmd_svc.GitCmdSvcInterface
	cli_pkg      cli_pkg.CliPkgInterface
	selectedMode string
}

func NewApp(
	os utils_pkg.OsInterface,
	check_svc svc.CheckSvcInterface,
	git_cmd_svc git_cmd_svc.GitCmdSvcInterface,
	cli_pkg cli_pkg.CliPkgInterface,
) *App {
	return &App{
		os:          os,
		check:       check_svc,
		git_cmd_svc: git_cmd_svc,
		cli_pkg:     cli_pkg,
	}
}

// 起動可能か確認
func (a *App) Check() {
	if err := a.check.CanStart(); err != nil {
		fmt.Println(err.Error())
		a.os.Exit(1)
	}
}

// 初期化処理
func (a *App) Init() {
	modes := config.ModeMenuConfig
	mode := flag.String("mode", "", "Mode: add or delete")
	flag.Parse()
	a.selectedMode = *mode
	if a.selectedMode == "" {
		a.selectedMode = a.cli_pkg.SelectFromMap("モードを選んでください", modes)
	}
}

// 実行処理
func (a *App) Run() {

	var err error
	var exitCode int
	switch a.selectedMode {
	case "git_com":
		exitCode, err = a.git_cmd_svc.GitCom()

	case "git_fork":
		exitCode, err = a.git_cmd_svc.GitFork()

	case "git_pr":
		exitCode, err = a.git_cmd_svc.GitPr()

	case "git_commit":
		exitCode, err = a.git_cmd_svc.GitCommit()

	case "git_commit_redo":
		exitCode, err = a.git_cmd_svc.GitCommitRedo()

	case "git_push":
		exitCode, err = a.git_cmd_svc.GitPush()

	case "git_pull":
		exitCode, err = a.git_cmd_svc.GitPull()

	case "git_fetch":
		exitCode, err = a.git_cmd_svc.GitFetch()
	case "exit":
		fmt.Println("👋 ツールを終了します。")
		a.os.Exit(0)

	default:
		fmt.Println("😕 未知のモードです:", a.selectedMode)
		a.os.Exit(0)
	}

	if err != nil {
		fmt.Println(err.Error())
		a.os.Exit(exitCode)
	}
}
