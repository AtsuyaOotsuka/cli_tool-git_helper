package config

import (
	"github.com/atylab-libs/go_cli_tool-libs/pkg/cli_pkg"
)

type MenuItem struct {
	Key   string
	Label string
}

var ModeMenuConfig = []cli_pkg.MenuItem{
	{Key: "git_pull", Label: "現在のブランチをpull"},
	{Key: "git_fork", Label: "現在のブランチからフォークして新しいブランチを作成"},
	{Key: "git_pr", Label: "プルリクエストを作成"},
	{Key: "git_commit", Label: "コミット"},
	{Key: "git_commit_redo", Label: "コミットを取り消す"},
	{Key: "git_push", Label: "プッシュ"},
	{Key: "git_fetch", Label: "リモートブランチをフェッチ"},
	{Key: "git_com", Label: "ディスクリプションを登録"},
	{Key: "exit", Label: "終了"}, // ← 最後に配置
}
