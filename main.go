package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/app"
	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/svc"
	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/svc/git_cmd_svc"

	"github.com/atylab-libs/go_cli_tool-libs/pkg/cli_pkg"
	"github.com/atylab-libs/go_cli_tool-libs/pkg/utils_pkg"
)

func main() {

	// シグナル通知チャネル作成
	sigs := make(chan os.Signal, 1)

	// 割り込みシグナルをキャッチ
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// 非同期で待機
	go func() {
		<-sigs
		fmt.Println("\n🚪 終了しました。")
		os.Exit(0)
	}()

	os := utils_pkg.NewOS()
	cli_pkg := cli_pkg.NewCliPkg()

	check_svc := svc.NewCheckService(os)
	git_cmd_svc := git_cmd_svc.NewGitCmdService(os, cli_pkg)

	app := app.NewApp(
		os,
		check_svc,
		git_cmd_svc,
		cli_pkg,
	)
	//app.Check()
	app.Init()
	app.Run()
}
