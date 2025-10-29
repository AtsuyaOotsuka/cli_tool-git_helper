package app

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/config"
	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/svc"
	"github.com/AtsuyaOotsuka/cli_tool-git_helper/internal/svc/git_cmd_svc"
	"github.com/AtsuyaOotsuka/cli_tool-git_helper/tests/mock/svc_mock"

	"github.com/atylab-libs/go_cli_tool-libs/pkg/cli_pkg"
	"github.com/atylab-libs/go_cli_tool-libs/pkg/utils_pkg"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/cli_pkg_mock"
	"github.com/atylab-libs/go_cli_tool-libs/tests/mock/pkg/utils_pkg_mock"

	"github.com/stretchr/testify/assert"
)

type AppTestStruct struct {
	os_mock          utils_pkg.OsInterface
	check_svc_mock   svc.CheckSvcInterface
	git_cmd_svc_mock git_cmd_svc.GitCmdSvcInterface
	cli_pkg_mock     cli_pkg.CliPkgInterface
}

func setUp() *AppTestStruct {
	os_mock := &utils_pkg_mock.OsMock{}
	check_svc_mock := &svc_mock.CheckSvcMock{}
	git_cmd_svc_mock := &svc_mock.GitCmdSvcMock{}
	cli_pkg_mock := &cli_pkg_mock.CliPkgMock{}

	return &AppTestStruct{
		os_mock:          os_mock,
		check_svc_mock:   check_svc_mock,
		git_cmd_svc_mock: git_cmd_svc_mock,
		cli_pkg_mock:     cli_pkg_mock,
	}
}

func createApp(a *AppTestStruct) *App {
	return NewApp(
		a.os_mock,
		a.check_svc_mock,
		a.git_cmd_svc_mock,
		a.cli_pkg_mock,
	)
}

func TestCheck(t *testing.T) {

	mock := setUp()

	check_svc_mock := mock.check_svc_mock.(*svc_mock.CheckSvcMock)
	check_svc_mock.On("CanStart").Return(nil)
	mock.check_svc_mock = check_svc_mock

	app := createApp(mock)
	app.Check()
}

func TestCheckWithErr(t *testing.T) {

	mock := setUp()

	check_svc_mock := mock.check_svc_mock.(*svc_mock.CheckSvcMock)
	check_svc_mock.On("CanStart").Return(fmt.Errorf("error"))
	mock.check_svc_mock = check_svc_mock

	os_mock := mock.os_mock.(*utils_pkg_mock.OsMock)
	os_mock.On("Exit", 2).Return()
	mock.os_mock = os_mock

	app := createApp(mock)
	app.Check()
}

func TestInit(t *testing.T) {
	mock := setUp()

	cli_pkg_mock := mock.cli_pkg_mock.(*cli_pkg_mock.CliPkgMock)
	cli_pkg_mock.On("SelectFromMap", "モードを選んでください", config.ModeMenuConfig).Return("git_com")

	app := createApp(mock)
	app.Init()
	assert.Equal(t, "git_com", app.selectedMode)
}

func TestInitWithFlag(t *testing.T) {

	mock := setUp()

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"dummy", "--mode=git_com"}

	app := createApp(mock)
	app.Init()

	assert.Equal(t, "git_com", app.selectedMode)

	t.Cleanup(func() {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	})
}

func TestRun(t *testing.T) {
	tests := []struct {
		name         string
		selectedMode string
		funcName     string
	}{
		{
			name:         "git_comモード",
			selectedMode: "git_com",
			funcName:     "GitCom",
		},
		{
			name:         "git_forkモード",
			selectedMode: "git_fork",
			funcName:     "GitFork",
		},
		{
			name:         "git_prモード",
			selectedMode: "git_pr",
			funcName:     "GitPr",
		},
		{
			name:         "git_commitモード",
			selectedMode: "git_commit",
			funcName:     "GitCommit",
		},
		{
			name:         "git_commit_redoモード",
			selectedMode: "git_commit_redo",
			funcName:     "GitCommitRedo",
		},
		{
			name:         "git_pushモード",
			selectedMode: "git_push",
			funcName:     "GitPush",
		},
		{
			name:         "git_pullモード",
			selectedMode: "git_pull",
			funcName:     "GitPull",
		},
		{
			name:         "git_fetchモード",
			selectedMode: "git_fetch",
			funcName:     "GitFetch",
		},
		{
			name:         "exitモード",
			selectedMode: "exit",
			funcName:     "",
		},
		{
			name:         "未知のモード",
			selectedMode: "unknown_mode",
			funcName:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := setUp()

			git_cmd_svc_mock := mock.git_cmd_svc_mock.(*svc_mock.GitCmdSvcMock)
			os_mock := mock.os_mock.(*utils_pkg_mock.OsMock)
			if tt.funcName != "" {
				git_cmd_svc_mock.On(tt.funcName).Return(0, nil)
			}
			os_mock.On("Exit", 0).Return()
			mock.git_cmd_svc_mock = git_cmd_svc_mock
			mock.os_mock = os_mock

			app := createApp(mock)
			app.selectedMode = tt.selectedMode
			app.Run()
		})

		t.Run(tt.name+"_with_error", func(t *testing.T) {
			mock := setUp()

			git_cmd_svc_mock := mock.git_cmd_svc_mock.(*svc_mock.GitCmdSvcMock)
			os_mock := mock.os_mock.(*utils_pkg_mock.OsMock)
			if tt.funcName != "" {
				git_cmd_svc_mock.On(tt.funcName).Return(1, fmt.Errorf("error"))
				os_mock.On("Exit", 1).Return()
			} else {
				os_mock.On("Exit", 0).Return()
			}
			mock.git_cmd_svc_mock = git_cmd_svc_mock
			mock.os_mock = os_mock

			app := createApp(mock)
			app.selectedMode = tt.selectedMode
			app.Run()
		})
	}
}
