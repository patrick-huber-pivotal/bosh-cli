// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-init/templatescompiler (interfaces: JobRenderer,JobListRenderer,RenderedJob,RenderedJobList,RenderedJobListArchive,RenderedJobListCompressor)

package mocks

import (
	property "github.com/cloudfoundry/bosh-init/common/property"
	gomock "github.com/cloudfoundry/bosh-init/internal/code.google.com/p/gomock/gomock"
	job "github.com/cloudfoundry/bosh-init/release/job"
	templatescompiler "github.com/cloudfoundry/bosh-init/templatescompiler"
)

// Mock of JobRenderer interface
type MockJobRenderer struct {
	ctrl     *gomock.Controller
	recorder *_MockJobRendererRecorder
}

// Recorder for MockJobRenderer (not exported)
type _MockJobRendererRecorder struct {
	mock *MockJobRenderer
}

func NewMockJobRenderer(ctrl *gomock.Controller) *MockJobRenderer {
	mock := &MockJobRenderer{ctrl: ctrl}
	mock.recorder = &_MockJobRendererRecorder{mock}
	return mock
}

func (_m *MockJobRenderer) EXPECT() *_MockJobRendererRecorder {
	return _m.recorder
}

func (_m *MockJobRenderer) Render(_param0 job.Job, _param1 property.Map, _param2 property.Map, _param3 string) (templatescompiler.RenderedJob, error) {
	ret := _m.ctrl.Call(_m, "Render", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(templatescompiler.RenderedJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobRendererRecorder) Render(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Render", arg0, arg1, arg2, arg3)
}

// Mock of JobListRenderer interface
type MockJobListRenderer struct {
	ctrl     *gomock.Controller
	recorder *_MockJobListRendererRecorder
}

// Recorder for MockJobListRenderer (not exported)
type _MockJobListRendererRecorder struct {
	mock *MockJobListRenderer
}

func NewMockJobListRenderer(ctrl *gomock.Controller) *MockJobListRenderer {
	mock := &MockJobListRenderer{ctrl: ctrl}
	mock.recorder = &_MockJobListRendererRecorder{mock}
	return mock
}

func (_m *MockJobListRenderer) EXPECT() *_MockJobListRendererRecorder {
	return _m.recorder
}

func (_m *MockJobListRenderer) Render(_param0 []job.Job, _param1 property.Map, _param2 property.Map, _param3 string) (templatescompiler.RenderedJobList, error) {
	ret := _m.ctrl.Call(_m, "Render", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(templatescompiler.RenderedJobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobListRendererRecorder) Render(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Render", arg0, arg1, arg2, arg3)
}

// Mock of RenderedJob interface
type MockRenderedJob struct {
	ctrl     *gomock.Controller
	recorder *_MockRenderedJobRecorder
}

// Recorder for MockRenderedJob (not exported)
type _MockRenderedJobRecorder struct {
	mock *MockRenderedJob
}

func NewMockRenderedJob(ctrl *gomock.Controller) *MockRenderedJob {
	mock := &MockRenderedJob{ctrl: ctrl}
	mock.recorder = &_MockRenderedJobRecorder{mock}
	return mock
}

func (_m *MockRenderedJob) EXPECT() *_MockRenderedJobRecorder {
	return _m.recorder
}

func (_m *MockRenderedJob) Delete() error {
	ret := _m.ctrl.Call(_m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRenderedJobRecorder) Delete() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete")
}

func (_m *MockRenderedJob) DeleteSilently() {
	_m.ctrl.Call(_m, "DeleteSilently")
}

func (_mr *_MockRenderedJobRecorder) DeleteSilently() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSilently")
}

func (_m *MockRenderedJob) Job() job.Job {
	ret := _m.ctrl.Call(_m, "Job")
	ret0, _ := ret[0].(job.Job)
	return ret0
}

func (_mr *_MockRenderedJobRecorder) Job() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Job")
}

func (_m *MockRenderedJob) Path() string {
	ret := _m.ctrl.Call(_m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRenderedJobRecorder) Path() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Path")
}

// Mock of RenderedJobList interface
type MockRenderedJobList struct {
	ctrl     *gomock.Controller
	recorder *_MockRenderedJobListRecorder
}

// Recorder for MockRenderedJobList (not exported)
type _MockRenderedJobListRecorder struct {
	mock *MockRenderedJobList
}

func NewMockRenderedJobList(ctrl *gomock.Controller) *MockRenderedJobList {
	mock := &MockRenderedJobList{ctrl: ctrl}
	mock.recorder = &_MockRenderedJobListRecorder{mock}
	return mock
}

func (_m *MockRenderedJobList) EXPECT() *_MockRenderedJobListRecorder {
	return _m.recorder
}

func (_m *MockRenderedJobList) Add(_param0 templatescompiler.RenderedJob) {
	_m.ctrl.Call(_m, "Add", _param0)
}

func (_mr *_MockRenderedJobListRecorder) Add(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0)
}

func (_m *MockRenderedJobList) All() []templatescompiler.RenderedJob {
	ret := _m.ctrl.Call(_m, "All")
	ret0, _ := ret[0].([]templatescompiler.RenderedJob)
	return ret0
}

func (_mr *_MockRenderedJobListRecorder) All() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "All")
}

func (_m *MockRenderedJobList) Delete() error {
	ret := _m.ctrl.Call(_m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRenderedJobListRecorder) Delete() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete")
}

func (_m *MockRenderedJobList) DeleteSilently() {
	_m.ctrl.Call(_m, "DeleteSilently")
}

func (_mr *_MockRenderedJobListRecorder) DeleteSilently() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSilently")
}

// Mock of RenderedJobListArchive interface
type MockRenderedJobListArchive struct {
	ctrl     *gomock.Controller
	recorder *_MockRenderedJobListArchiveRecorder
}

// Recorder for MockRenderedJobListArchive (not exported)
type _MockRenderedJobListArchiveRecorder struct {
	mock *MockRenderedJobListArchive
}

func NewMockRenderedJobListArchive(ctrl *gomock.Controller) *MockRenderedJobListArchive {
	mock := &MockRenderedJobListArchive{ctrl: ctrl}
	mock.recorder = &_MockRenderedJobListArchiveRecorder{mock}
	return mock
}

func (_m *MockRenderedJobListArchive) EXPECT() *_MockRenderedJobListArchiveRecorder {
	return _m.recorder
}

func (_m *MockRenderedJobListArchive) Delete() error {
	ret := _m.ctrl.Call(_m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRenderedJobListArchiveRecorder) Delete() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete")
}

func (_m *MockRenderedJobListArchive) DeleteSilently() {
	_m.ctrl.Call(_m, "DeleteSilently")
}

func (_mr *_MockRenderedJobListArchiveRecorder) DeleteSilently() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSilently")
}

func (_m *MockRenderedJobListArchive) Fingerprint() string {
	ret := _m.ctrl.Call(_m, "Fingerprint")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRenderedJobListArchiveRecorder) Fingerprint() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fingerprint")
}

func (_m *MockRenderedJobListArchive) List() templatescompiler.RenderedJobList {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].(templatescompiler.RenderedJobList)
	return ret0
}

func (_mr *_MockRenderedJobListArchiveRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockRenderedJobListArchive) Path() string {
	ret := _m.ctrl.Call(_m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRenderedJobListArchiveRecorder) Path() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Path")
}

func (_m *MockRenderedJobListArchive) SHA1() string {
	ret := _m.ctrl.Call(_m, "SHA1")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRenderedJobListArchiveRecorder) SHA1() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SHA1")
}

// Mock of RenderedJobListCompressor interface
type MockRenderedJobListCompressor struct {
	ctrl     *gomock.Controller
	recorder *_MockRenderedJobListCompressorRecorder
}

// Recorder for MockRenderedJobListCompressor (not exported)
type _MockRenderedJobListCompressorRecorder struct {
	mock *MockRenderedJobListCompressor
}

func NewMockRenderedJobListCompressor(ctrl *gomock.Controller) *MockRenderedJobListCompressor {
	mock := &MockRenderedJobListCompressor{ctrl: ctrl}
	mock.recorder = &_MockRenderedJobListCompressorRecorder{mock}
	return mock
}

func (_m *MockRenderedJobListCompressor) EXPECT() *_MockRenderedJobListCompressorRecorder {
	return _m.recorder
}

func (_m *MockRenderedJobListCompressor) Compress(_param0 templatescompiler.RenderedJobList) (templatescompiler.RenderedJobListArchive, error) {
	ret := _m.ctrl.Call(_m, "Compress", _param0)
	ret0, _ := ret[0].(templatescompiler.RenderedJobListArchive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRenderedJobListCompressorRecorder) Compress(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Compress", arg0)
}
