package git_cmd_svc

import (
	"fmt"
	"strings"
)

func (g *GitCmdSvc) getParent(
	currentBranch string,
	repoName string,
) (string, error) {
	baseDir, err := g.os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/.git_project/%s/%s.parent", baseDir, repoName, currentBranch)

	if _, err := g.os.Stat(path); g.os.IsNotExist(err) {
		return "", fmt.Errorf("親ブランチ情報が存在しません")
	}

	content, readFileErr := g.os.ReadFile(path)
	if readFileErr != nil {
		return "", fmt.Errorf("親ブランチ情報の読み込みに失敗しました: %w", readFileErr)
	}

	text := string(content)
	text = strings.TrimSpace(text)
	if text == "" {
		return "", fmt.Errorf("親ブランチ情報が空です")
	}
	return text, nil
}

func (g *GitCmdSvc) GitPr() (int, error) {
	currentBranch, err := g.getCurrentBranchName()
	if err != nil {
		return 1, fmt.Errorf("カレントブランチ名が取得できませんでした")
	}

	repoName, err := g.getRepoName()
	if err != nil {
		return 1, fmt.Errorf("リポジトリ名が取得できませんでした")
	}

	parentBranch, err := g.getParent(currentBranch, repoName)
	if err != nil {
		return 1, fmt.Errorf("親ブランチ情報の取得に失敗しました: %w", err)
	}

	repoBaseUrl, getRepoBaseUrlErr := g.getRemoteUrl()
	if getRepoBaseUrlErr != nil {
		return 1, fmt.Errorf("リモートリポジトリのURL取得に失敗しました: %w", getRepoBaseUrlErr)
	}
	url := fmt.Sprintf(repoBaseUrl+"/compare/%s...%s", parentBranch, currentBranch)
	fmt.Println("URLを開きます:", url)

	err = g.os.CommandRun("open", url)
	if err != nil {
		return 1, fmt.Errorf("ブラウザの起動に失敗しました: %w", err)
	}

	return 0, nil
}
