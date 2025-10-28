package git_cmd_svc

import (
	"fmt"
)

func (g *GitCmdSvc) GitFetch() (int, error) {
	targetBranch, err := g.cli_pkg.Read("フェッチするブランチ名を入力してください: ")

	if err != nil {
		if err.Error() == "EOF" {
			return 0, nil
		}
		return 1, err
	}

	if targetBranch == "" {
		return 1, fmt.Errorf("ブランチ名が空です")
	}

	fmt.Printf("ブランチ '%s' をフェッチします。\n", targetBranch)
	result, err := g.cli_pkg.Confirm("この操作を実行しますか？")
	if err != nil {
		return 1, err
	}
	if !result {
		return 0, fmt.Errorf("操作がキャンセルされました。")
	}
	cmd, err := g.os.Command("git", "fetch", "origin", targetBranch)
	if err != nil {
		return 1, err
	}
	fmt.Printf("✅ フェッチが成功しました！ \n%s", string(cmd))
	return 0, nil
}
