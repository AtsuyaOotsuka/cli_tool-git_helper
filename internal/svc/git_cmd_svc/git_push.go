package git_cmd_svc

import "fmt"

func (g *GitCmdSvc) GitPush() (int, error) {
	currentBranch, err := g.getCurrentBranchName()
	if err != nil {
		return 1, fmt.Errorf("カレントブランチ名が取得できませんでした")
	}

	result, err := g.cli_pkg.Confirm(fmt.Sprintf("ブランチ %s を push しますか？", currentBranch))
	if err != nil {
		return 1, err
	}
	if !result {
		return 0, fmt.Errorf("操作がキャンセルされました。")
	}

	cmd, err := g.os.Command("git", "push", "origin", currentBranch)
	if err != nil {
		return 1, err
	}
	fmt.Printf("✅ Push が成功しました！ \n%s", string(cmd))
	return 0, nil
}
