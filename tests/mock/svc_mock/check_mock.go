package svc_mock

import "github.com/stretchr/testify/mock"

type CheckSvcMock struct {
	mock.Mock
}

func (csm *CheckSvcMock) CanStart() error {
	args := csm.Called()
	return args.Error(0)
}
