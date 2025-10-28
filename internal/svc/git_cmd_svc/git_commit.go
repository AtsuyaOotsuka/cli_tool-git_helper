package git_cmd_svc

import "fmt"

func (g *GitCmdSvc) GitCommit() (int, error) {

	currentBranch, err := g.getCurrentBranchName()
	if err != nil {
		return 1, fmt.Errorf("カレントブランチ名が取得できませんでした")
	}
	prefix := g.getPrefix(currentBranch)

	fmt.Printf("現在のブランチ: %s\n", currentBranch)
	fmt.Printf("プレフィックス: %s\n", prefix)

	message, err := g.cli_pkg.Read("コミットメッセージを入力してください: ")
	if err != nil {
		if err.Error() == "EOF" {
			return 0, nil
		}
		return 1, err
	}

	if message == "" {
		return 1, fmt.Errorf("コミットメッセージが空です")
	}
	if prefix != "" {
		message = prefix + " " + message
	}
	cmd, err := g.os.Command("git", "commit", "-m", message)
	if err != nil {
		return 1, fmt.Errorf("git コマンドの実行中にエラーが発生しました: %v", err)
	}

	fmt.Printf("✅ コミットが成功しました！ \n%s", string(cmd))
	return 0, nil
}
