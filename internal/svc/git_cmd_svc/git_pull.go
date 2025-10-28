package git_cmd_svc

import "fmt"

func (g *GitCmdSvc) GitPull() (int, error) {
	currentBranch, err := g.getCurrentBranchName()
	if err != nil {
		return 1, fmt.Errorf("カレントブランチ名が取得できませんでした")
	}

	result, err := g.cli_pkg.Confirm(fmt.Sprintf("ブランチ %s を pull しますか？", currentBranch))
	if err != nil {
		return 1, err
	}
	if !result {
		return 0, fmt.Errorf("操作がキャンセルされました。")
	}

	cmd, err := g.os.Command("git", "pull", "origin", currentBranch)
	if err != nil {
		return 1, err
	}
	fmt.Printf("✅ Pull が成功しました！ \n%s", string(cmd))

	return 0, nil
}
