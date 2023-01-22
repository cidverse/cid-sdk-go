// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	cidsdk "github.com/cidverse/cid-sdk-go"
	mock "github.com/stretchr/testify/mock"
)

// SDKClient is an autogenerated mock type for the SDKClient type
type SDKClient struct {
	mock.Mock
}

// ArtifactDownload provides a mock function with given fields: request
func (_m *SDKClient) ArtifactDownload(request cidsdk.ArtifactDownloadRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(cidsdk.ArtifactDownloadRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ArtifactList provides a mock function with given fields: request
func (_m *SDKClient) ArtifactList(request cidsdk.ArtifactListRequest) (*[]cidsdk.ActionArtifact, error) {
	ret := _m.Called(request)

	var r0 *[]cidsdk.ActionArtifact
	if rf, ok := ret.Get(0).(func(cidsdk.ArtifactListRequest) *[]cidsdk.ActionArtifact); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]cidsdk.ActionArtifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cidsdk.ArtifactListRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ArtifactUpload provides a mock function with given fields: request
func (_m *SDKClient) ArtifactUpload(request cidsdk.ArtifactUploadRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(cidsdk.ArtifactUploadRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CurrentConfig provides a mock function with given fields:
func (_m *SDKClient) CurrentConfig() (*cidsdk.CurrentConfig, error) {
	ret := _m.Called()

	var r0 *cidsdk.CurrentConfig
	if rf, ok := ret.Get(0).(func() *cidsdk.CurrentConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cidsdk.CurrentConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentModule provides a mock function with given fields:
func (_m *SDKClient) CurrentModule() (*cidsdk.ProjectModule, error) {
	ret := _m.Called()

	var r0 *cidsdk.ProjectModule
	if rf, ok := ret.Get(0).(func() *cidsdk.ProjectModule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cidsdk.ProjectModule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Env provides a mock function with given fields:
func (_m *SDKClient) Env() (map[string]string, error) {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteCommand provides a mock function with given fields: req
func (_m *SDKClient) ExecuteCommand(req cidsdk.ExecuteCommandRequest) (*cidsdk.ExecuteCommandResponse, error) {
	ret := _m.Called(req)

	var r0 *cidsdk.ExecuteCommandResponse
	if rf, ok := ret.Get(0).(func(cidsdk.ExecuteCommandRequest) *cidsdk.ExecuteCommandResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cidsdk.ExecuteCommandResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cidsdk.ExecuteCommandRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileList provides a mock function with given fields: req
func (_m *SDKClient) FileList(req cidsdk.FileRequest) ([]cidsdk.File, error) {
	ret := _m.Called(req)

	var r0 []cidsdk.File
	if rf, ok := ret.Get(0).(func(cidsdk.FileRequest) []cidsdk.File); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cidsdk.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cidsdk.FileRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileRead provides a mock function with given fields: file
func (_m *SDKClient) FileRead(file string) (string, error) {
	ret := _m.Called(file)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileRemove provides a mock function with given fields: file
func (_m *SDKClient) FileRemove(file string) error {
	ret := _m.Called(file)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileRename provides a mock function with given fields: old, new
func (_m *SDKClient) FileRename(old string, new string) error {
	ret := _m.Called(old, new)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(old, new)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FileWrite provides a mock function with given fields: file, content
func (_m *SDKClient) FileWrite(file string, content []byte) error {
	ret := _m.Called(file, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(file, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Health provides a mock function with given fields:
func (_m *SDKClient) Health() (*cidsdk.HealthcheckResponse, error) {
	ret := _m.Called()

	var r0 *cidsdk.HealthcheckResponse
	if rf, ok := ret.Get(0).(func() *cidsdk.HealthcheckResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cidsdk.HealthcheckResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Log provides a mock function with given fields: req
func (_m *SDKClient) Log(req cidsdk.LogMessageRequest) error {
	return nil
}

// ModuleAction provides a mock function with given fields: cfg
func (_m *SDKClient) ModuleAction(cfg interface{}) (cidsdk.ModuleActionData, error) {
	ret := _m.Called(cfg)

	var r0 cidsdk.ModuleActionData
	if rf, ok := ret.Get(0).(func(interface{}) cidsdk.ModuleActionData); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Get(0).(cidsdk.ModuleActionData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modules provides a mock function with given fields:
func (_m *SDKClient) Modules() (*[]cidsdk.ProjectModule, error) {
	ret := _m.Called()

	var r0 *[]cidsdk.ProjectModule
	if rf, ok := ret.Get(0).(func() *[]cidsdk.ProjectModule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]cidsdk.ProjectModule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectAction provides a mock function with given fields: cfg
func (_m *SDKClient) ProjectAction(cfg interface{}) (cidsdk.ProjectActionData, error) {
	ret := _m.Called(cfg)

	var r0 cidsdk.ProjectActionData
	if rf, ok := ret.Get(0).(func(interface{}) cidsdk.ProjectActionData); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Get(0).(cidsdk.ProjectActionData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UUID provides a mock function with given fields:
func (_m *SDKClient) UUID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VCSCommitByHash provides a mock function with given fields: request
func (_m *SDKClient) VCSCommitByHash(request cidsdk.VCSCommitByHashRequest) (*cidsdk.VCSCommit, error) {
	ret := _m.Called(request)

	var r0 *cidsdk.VCSCommit
	if rf, ok := ret.Get(0).(func(cidsdk.VCSCommitByHashRequest) *cidsdk.VCSCommit); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cidsdk.VCSCommit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cidsdk.VCSCommitByHashRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VCSCommits provides a mock function with given fields: request
func (_m *SDKClient) VCSCommits(request cidsdk.VCSCommitsRequest) (*[]cidsdk.VCSCommit, error) {
	ret := _m.Called(request)

	var r0 *[]cidsdk.VCSCommit
	if rf, ok := ret.Get(0).(func(cidsdk.VCSCommitsRequest) *[]cidsdk.VCSCommit); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]cidsdk.VCSCommit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cidsdk.VCSCommitsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VCSReleases provides a mock function with given fields:
func (_m *SDKClient) VCSReleases() (*[]cidsdk.VCSRelease, error) {
	ret := _m.Called()

	var r0 *[]cidsdk.VCSRelease
	if rf, ok := ret.Get(0).(func() *[]cidsdk.VCSRelease); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]cidsdk.VCSRelease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VCSTags provides a mock function with given fields:
func (_m *SDKClient) VCSTags() (*[]cidsdk.VCSTag, error) {
	ret := _m.Called()

	var r0 *[]cidsdk.VCSTag
	if rf, ok := ret.Get(0).(func() *[]cidsdk.VCSTag); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]cidsdk.VCSTag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSDKClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewSDKClient creates a new instance of SDKClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSDKClient(t mockConstructorTestingTNewSDKClient) *SDKClient {
	mock := &SDKClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
