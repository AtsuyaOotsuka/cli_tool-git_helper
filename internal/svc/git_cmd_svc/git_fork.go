package git_cmd_svc

import (
	"fmt"
	"os"
	"path/filepath"
)

func (g *GitCmdSvc) setParent(
	newBranch string,
	parentBranch string,
	repoName string,
) error {
	baseDir, err := g.os.UserHomeDir()
	if err != nil {
		return err
	}
	dirPath := filepath.Join(baseDir, ".git_project", repoName)

	if err := g.os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("ディレクトリ作成に失敗: %w", err)
	}

	filePath := filepath.Join(dirPath, fmt.Sprintf("%s.parent", newBranch))

	if err := g.os.WriteFile(filePath, []byte(parentBranch), 0644); err != nil {
		return fmt.Errorf("ファイル書き込みに失敗: %w", err)
	}

	return nil
}

func (g *GitCmdSvc) GitFork() (int, error) {
	var err error
	currentBranch, err := g.getCurrentBranchName()
	if err != nil {
		return 1, fmt.Errorf("カレントブランチ名が取得できませんでした")
	}
	repoName, err := g.getRepoName()
	if err != nil {
		return 1, fmt.Errorf("リポジトリ名が取得できませんでした")
	}

	fmt.Printf("現在のブランチ: %s\n", currentBranch)
	fmt.Printf("リポジトリ名: %s\n", repoName)

	newBranch, err := g.cli_pkg.Read("新しいブランチ名を入力してください: ")
	if err != nil {
		if err.Error() == "EOF" {
			return 0, nil
		}
		return 1, err
	}

	if newBranch == currentBranch {
		return 1, fmt.Errorf("新しいブランチ名は現在のブランチ名と異なる必要があります。")
	}

	fmt.Printf("新しいブランチ '%s' を作成します。\n", newBranch)
	result, err := g.cli_pkg.Confirm("この操作を実行しますか？")
	if err != nil {
		return 1, err
	}
	if !result {
		return 0, fmt.Errorf("操作がキャンセルされました。")
	}
	fmt.Printf("新しいブランチ '%s' を作成します...\n", newBranch)
	_, err = g.os.Command("git", "checkout", "-b", newBranch)
	if err != nil {
		return 1, fmt.Errorf("git コマンドの実行中にエラーが発生しました: %v", err)
	}

	err = g.setParent(newBranch, currentBranch, repoName)
	if err != nil {
		return 1, fmt.Errorf("親ブランチの設定中にエラーが発生しました: %v", err)
	}

	fmt.Printf("✅ ブランチ '%s' の作成が成功しました！\n", newBranch)

	return 0, nil
}
