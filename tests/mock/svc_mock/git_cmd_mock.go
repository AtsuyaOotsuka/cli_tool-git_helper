package svc_mock

import (
	"github.com/stretchr/testify/mock"
)

type GitCmdSvcMock struct {
	mock.Mock
}

func (m *GitCmdSvcMock) GitCom() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *GitCmdSvcMock) GitFork() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *GitCmdSvcMock) GitPr() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *GitCmdSvcMock) GitCommit() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *GitCmdSvcMock) GitCommitRedo() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *GitCmdSvcMock) GitPush() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *GitCmdSvcMock) GitPull() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *GitCmdSvcMock) GitFetch() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
