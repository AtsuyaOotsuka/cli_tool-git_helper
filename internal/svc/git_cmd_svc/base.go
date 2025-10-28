package git_cmd_svc

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/atylab-libs/go_cli_tool-libs/pkg/cli_pkg"
	"github.com/atylab-libs/go_cli_tool-libs/pkg/utils_pkg"
)

type GitCmdSvcInterface interface {
	GitFetch() (int, error)
	GitPull() (int, error)
	GitPush() (int, error)
	GitCommit() (int, error)
	GitCommitRedo() (int, error)
	GitPr() (int, error)
	GitFork() (int, error)
	GitCom() (int, error)
}

type GitCmdSvc struct {
	os      utils_pkg.OsInterface
	cli_pkg cli_pkg.CliPkgInterface
}

func NewGitCmdService(
	os utils_pkg.OsInterface,
	cli_pkg cli_pkg.CliPkgInterface,
) GitCmdSvcInterface {
	return &GitCmdSvc{
		os:      os,
		cli_pkg: cli_pkg,
	}
}

func (g *GitCmdSvc) getCurrentBranchName() (string, error) {
	out, err := g.os.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func (g *GitCmdSvc) getPrefix(branchName string) string {
	re := regexp.MustCompile(`^[^-]+-[0-9]+`)
	match := re.FindString(branchName)
	return match
}

func (g *GitCmdSvc) getRemoteUrl() (string, error) {
	out, err := g.os.Command("git", "config", "--get", "remote.origin.url")
	if err != nil {
		return "", err
	}

	url := strings.TrimSpace(string(out))
	// URLパターン： https://github.com/user/repo.git または git@github.com:user/repo.git
	re := regexp.MustCompile(`github\.com[:/](?P<user>[^/]+)/(?P<repo>[^/]+)(\.git)?$`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 3 {
		return "", fmt.Errorf("GitHubユーザー名・リポジトリ名の抽出に失敗しました")
	}

	user := matches[1]
	repo := matches[2]

	repo = strings.TrimSuffix(repo, ".git") // .gitを削除

	return fmt.Sprintf("https://github.com/%s/%s", user, repo), nil
}

func (g *GitCmdSvc) getRepoName() (string, error) {
	out, err := g.os.Command("git", "config", "--get", "remote.origin.url")
	if err != nil {
		return "", err
	}

	// 最後の部分を取る（例: "repo.git"）
	base := path.Base(strings.TrimSpace(string(out)))

	repo := strings.TrimSuffix(base, ".git")
	return repo, nil
}
