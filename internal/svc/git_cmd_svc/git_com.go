package git_cmd_svc

import (
	"fmt"
)

func (g *GitCmdSvc) GitCom() (int, error) {
	currentBranch, err := g.getCurrentBranchName()
	if err != nil {
		return 1, fmt.Errorf("カレントブランチ名が取得できませんでした")
	}

	description, err := g.cli_pkg.Read("ディスクリプションを入力してください: ")
	if err != nil {
		if err.Error() == "EOF" {
			return 0, nil
		}
		return 1, err
	}

	fmt.Printf("設定中: ブランチ='%s', ディスクリプション='%s'", currentBranch, description)

	cmd, err := g.os.Command("git", "config", "branch."+currentBranch+".description", description)

	if err != nil {
		return 1, fmt.Errorf("git コマンドの実行中にエラーが発生しました: %v", err)
	}

	fmt.Printf("ブランチ '%s' にディスクリプションを設定しました \n%s", currentBranch, string(cmd))
	return 0, nil
}
