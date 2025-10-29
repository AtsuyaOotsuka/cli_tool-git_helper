package git_cmd_svc

import "fmt"

func (g *GitCmdSvc) GitCommitRedo() (int, error) {
	commitMessage, err := g.os.Command("git", "log", "-1", "--pretty=%s")
	if err != nil {
		return 1, fmt.Errorf("最新のコミットメッセージの取得に失敗しました: %v", err)
	}

	fmt.Printf("最終コミット: %s\n", commitMessage)
	result, err := g.cli_pkg.Confirm("このコミットを取り消しますか？")
	if err != nil {
		return 1, err
	}
	if !result {
		return 0, fmt.Errorf("操作がキャンセルされました。")
	}

	cmd, err := g.os.Command("git", "reset", "--soft", "HEAD~1")
	if err != nil {
		return 1, fmt.Errorf("コミットの取り消しに失敗しました: %v", err)
	}

	fmt.Printf("✅ コミットが取り消されました！ \n%s", string(cmd))
	return 0, nil
}
