// Code generated by counterfeiter. DO NOT EDIT.
package corefakes

import (
	"sync"

	"github.com/ccremer/greposync/core"
)

type FakePullRequest struct {
	GetBodyStub        func() string
	getBodyMutex       sync.RWMutex
	getBodyArgsForCall []struct {
	}
	getBodyReturns struct {
		result1 string
	}
	getBodyReturnsOnCall map[int]struct {
		result1 string
	}
	GetCommitBranchStub        func() string
	getCommitBranchMutex       sync.RWMutex
	getCommitBranchArgsForCall []struct {
	}
	getCommitBranchReturns struct {
		result1 string
	}
	getCommitBranchReturnsOnCall map[int]struct {
		result1 string
	}
	GetLabelsStub        func() []core.Label
	getLabelsMutex       sync.RWMutex
	getLabelsArgsForCall []struct {
	}
	getLabelsReturns struct {
		result1 []core.Label
	}
	getLabelsReturnsOnCall map[int]struct {
		result1 []core.Label
	}
	GetTargetBranchStub        func() string
	getTargetBranchMutex       sync.RWMutex
	getTargetBranchArgsForCall []struct {
	}
	getTargetBranchReturns struct {
		result1 string
	}
	getTargetBranchReturnsOnCall map[int]struct {
		result1 string
	}
	GetTitleStub        func() string
	getTitleMutex       sync.RWMutex
	getTitleArgsForCall []struct {
	}
	getTitleReturns struct {
		result1 string
	}
	getTitleReturnsOnCall map[int]struct {
		result1 string
	}
	SetBodyStub        func(string)
	setBodyMutex       sync.RWMutex
	setBodyArgsForCall []struct {
		arg1 string
	}
	SetCommitBranchStub        func(string) string
	setCommitBranchMutex       sync.RWMutex
	setCommitBranchArgsForCall []struct {
		arg1 string
	}
	setCommitBranchReturns struct {
		result1 string
	}
	setCommitBranchReturnsOnCall map[int]struct {
		result1 string
	}
	SetLabelsStub        func([]core.Label)
	setLabelsMutex       sync.RWMutex
	setLabelsArgsForCall []struct {
		arg1 []core.Label
	}
	SetTargetBranchStub        func(string)
	setTargetBranchMutex       sync.RWMutex
	setTargetBranchArgsForCall []struct {
		arg1 string
	}
	SetTitleStub        func(string)
	setTitleMutex       sync.RWMutex
	setTitleArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePullRequest) GetBody() string {
	fake.getBodyMutex.Lock()
	ret, specificReturn := fake.getBodyReturnsOnCall[len(fake.getBodyArgsForCall)]
	fake.getBodyArgsForCall = append(fake.getBodyArgsForCall, struct {
	}{})
	stub := fake.GetBodyStub
	fakeReturns := fake.getBodyReturns
	fake.recordInvocation("GetBody", []interface{}{})
	fake.getBodyMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePullRequest) GetBodyCallCount() int {
	fake.getBodyMutex.RLock()
	defer fake.getBodyMutex.RUnlock()
	return len(fake.getBodyArgsForCall)
}

func (fake *FakePullRequest) GetBodyCalls(stub func() string) {
	fake.getBodyMutex.Lock()
	defer fake.getBodyMutex.Unlock()
	fake.GetBodyStub = stub
}

func (fake *FakePullRequest) GetBodyReturns(result1 string) {
	fake.getBodyMutex.Lock()
	defer fake.getBodyMutex.Unlock()
	fake.GetBodyStub = nil
	fake.getBodyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) GetBodyReturnsOnCall(i int, result1 string) {
	fake.getBodyMutex.Lock()
	defer fake.getBodyMutex.Unlock()
	fake.GetBodyStub = nil
	if fake.getBodyReturnsOnCall == nil {
		fake.getBodyReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getBodyReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) GetCommitBranch() string {
	fake.getCommitBranchMutex.Lock()
	ret, specificReturn := fake.getCommitBranchReturnsOnCall[len(fake.getCommitBranchArgsForCall)]
	fake.getCommitBranchArgsForCall = append(fake.getCommitBranchArgsForCall, struct {
	}{})
	stub := fake.GetCommitBranchStub
	fakeReturns := fake.getCommitBranchReturns
	fake.recordInvocation("GetCommitBranch", []interface{}{})
	fake.getCommitBranchMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePullRequest) GetCommitBranchCallCount() int {
	fake.getCommitBranchMutex.RLock()
	defer fake.getCommitBranchMutex.RUnlock()
	return len(fake.getCommitBranchArgsForCall)
}

func (fake *FakePullRequest) GetCommitBranchCalls(stub func() string) {
	fake.getCommitBranchMutex.Lock()
	defer fake.getCommitBranchMutex.Unlock()
	fake.GetCommitBranchStub = stub
}

func (fake *FakePullRequest) GetCommitBranchReturns(result1 string) {
	fake.getCommitBranchMutex.Lock()
	defer fake.getCommitBranchMutex.Unlock()
	fake.GetCommitBranchStub = nil
	fake.getCommitBranchReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) GetCommitBranchReturnsOnCall(i int, result1 string) {
	fake.getCommitBranchMutex.Lock()
	defer fake.getCommitBranchMutex.Unlock()
	fake.GetCommitBranchStub = nil
	if fake.getCommitBranchReturnsOnCall == nil {
		fake.getCommitBranchReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getCommitBranchReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) GetLabels() []core.Label {
	fake.getLabelsMutex.Lock()
	ret, specificReturn := fake.getLabelsReturnsOnCall[len(fake.getLabelsArgsForCall)]
	fake.getLabelsArgsForCall = append(fake.getLabelsArgsForCall, struct {
	}{})
	stub := fake.GetLabelsStub
	fakeReturns := fake.getLabelsReturns
	fake.recordInvocation("GetLabels", []interface{}{})
	fake.getLabelsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePullRequest) GetLabelsCallCount() int {
	fake.getLabelsMutex.RLock()
	defer fake.getLabelsMutex.RUnlock()
	return len(fake.getLabelsArgsForCall)
}

func (fake *FakePullRequest) GetLabelsCalls(stub func() []core.Label) {
	fake.getLabelsMutex.Lock()
	defer fake.getLabelsMutex.Unlock()
	fake.GetLabelsStub = stub
}

func (fake *FakePullRequest) GetLabelsReturns(result1 []core.Label) {
	fake.getLabelsMutex.Lock()
	defer fake.getLabelsMutex.Unlock()
	fake.GetLabelsStub = nil
	fake.getLabelsReturns = struct {
		result1 []core.Label
	}{result1}
}

func (fake *FakePullRequest) GetLabelsReturnsOnCall(i int, result1 []core.Label) {
	fake.getLabelsMutex.Lock()
	defer fake.getLabelsMutex.Unlock()
	fake.GetLabelsStub = nil
	if fake.getLabelsReturnsOnCall == nil {
		fake.getLabelsReturnsOnCall = make(map[int]struct {
			result1 []core.Label
		})
	}
	fake.getLabelsReturnsOnCall[i] = struct {
		result1 []core.Label
	}{result1}
}

func (fake *FakePullRequest) GetTargetBranch() string {
	fake.getTargetBranchMutex.Lock()
	ret, specificReturn := fake.getTargetBranchReturnsOnCall[len(fake.getTargetBranchArgsForCall)]
	fake.getTargetBranchArgsForCall = append(fake.getTargetBranchArgsForCall, struct {
	}{})
	stub := fake.GetTargetBranchStub
	fakeReturns := fake.getTargetBranchReturns
	fake.recordInvocation("GetTargetBranch", []interface{}{})
	fake.getTargetBranchMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePullRequest) GetTargetBranchCallCount() int {
	fake.getTargetBranchMutex.RLock()
	defer fake.getTargetBranchMutex.RUnlock()
	return len(fake.getTargetBranchArgsForCall)
}

func (fake *FakePullRequest) GetTargetBranchCalls(stub func() string) {
	fake.getTargetBranchMutex.Lock()
	defer fake.getTargetBranchMutex.Unlock()
	fake.GetTargetBranchStub = stub
}

func (fake *FakePullRequest) GetTargetBranchReturns(result1 string) {
	fake.getTargetBranchMutex.Lock()
	defer fake.getTargetBranchMutex.Unlock()
	fake.GetTargetBranchStub = nil
	fake.getTargetBranchReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) GetTargetBranchReturnsOnCall(i int, result1 string) {
	fake.getTargetBranchMutex.Lock()
	defer fake.getTargetBranchMutex.Unlock()
	fake.GetTargetBranchStub = nil
	if fake.getTargetBranchReturnsOnCall == nil {
		fake.getTargetBranchReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getTargetBranchReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) GetTitle() string {
	fake.getTitleMutex.Lock()
	ret, specificReturn := fake.getTitleReturnsOnCall[len(fake.getTitleArgsForCall)]
	fake.getTitleArgsForCall = append(fake.getTitleArgsForCall, struct {
	}{})
	stub := fake.GetTitleStub
	fakeReturns := fake.getTitleReturns
	fake.recordInvocation("GetTitle", []interface{}{})
	fake.getTitleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePullRequest) GetTitleCallCount() int {
	fake.getTitleMutex.RLock()
	defer fake.getTitleMutex.RUnlock()
	return len(fake.getTitleArgsForCall)
}

func (fake *FakePullRequest) GetTitleCalls(stub func() string) {
	fake.getTitleMutex.Lock()
	defer fake.getTitleMutex.Unlock()
	fake.GetTitleStub = stub
}

func (fake *FakePullRequest) GetTitleReturns(result1 string) {
	fake.getTitleMutex.Lock()
	defer fake.getTitleMutex.Unlock()
	fake.GetTitleStub = nil
	fake.getTitleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) GetTitleReturnsOnCall(i int, result1 string) {
	fake.getTitleMutex.Lock()
	defer fake.getTitleMutex.Unlock()
	fake.GetTitleStub = nil
	if fake.getTitleReturnsOnCall == nil {
		fake.getTitleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getTitleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) SetBody(arg1 string) {
	fake.setBodyMutex.Lock()
	fake.setBodyArgsForCall = append(fake.setBodyArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetBodyStub
	fake.recordInvocation("SetBody", []interface{}{arg1})
	fake.setBodyMutex.Unlock()
	if stub != nil {
		fake.SetBodyStub(arg1)
	}
}

func (fake *FakePullRequest) SetBodyCallCount() int {
	fake.setBodyMutex.RLock()
	defer fake.setBodyMutex.RUnlock()
	return len(fake.setBodyArgsForCall)
}

func (fake *FakePullRequest) SetBodyCalls(stub func(string)) {
	fake.setBodyMutex.Lock()
	defer fake.setBodyMutex.Unlock()
	fake.SetBodyStub = stub
}

func (fake *FakePullRequest) SetBodyArgsForCall(i int) string {
	fake.setBodyMutex.RLock()
	defer fake.setBodyMutex.RUnlock()
	argsForCall := fake.setBodyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePullRequest) SetCommitBranch(arg1 string) string {
	fake.setCommitBranchMutex.Lock()
	ret, specificReturn := fake.setCommitBranchReturnsOnCall[len(fake.setCommitBranchArgsForCall)]
	fake.setCommitBranchArgsForCall = append(fake.setCommitBranchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetCommitBranchStub
	fakeReturns := fake.setCommitBranchReturns
	fake.recordInvocation("SetCommitBranch", []interface{}{arg1})
	fake.setCommitBranchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePullRequest) SetCommitBranchCallCount() int {
	fake.setCommitBranchMutex.RLock()
	defer fake.setCommitBranchMutex.RUnlock()
	return len(fake.setCommitBranchArgsForCall)
}

func (fake *FakePullRequest) SetCommitBranchCalls(stub func(string) string) {
	fake.setCommitBranchMutex.Lock()
	defer fake.setCommitBranchMutex.Unlock()
	fake.SetCommitBranchStub = stub
}

func (fake *FakePullRequest) SetCommitBranchArgsForCall(i int) string {
	fake.setCommitBranchMutex.RLock()
	defer fake.setCommitBranchMutex.RUnlock()
	argsForCall := fake.setCommitBranchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePullRequest) SetCommitBranchReturns(result1 string) {
	fake.setCommitBranchMutex.Lock()
	defer fake.setCommitBranchMutex.Unlock()
	fake.SetCommitBranchStub = nil
	fake.setCommitBranchReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) SetCommitBranchReturnsOnCall(i int, result1 string) {
	fake.setCommitBranchMutex.Lock()
	defer fake.setCommitBranchMutex.Unlock()
	fake.SetCommitBranchStub = nil
	if fake.setCommitBranchReturnsOnCall == nil {
		fake.setCommitBranchReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.setCommitBranchReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequest) SetLabels(arg1 []core.Label) {
	var arg1Copy []core.Label
	if arg1 != nil {
		arg1Copy = make([]core.Label, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setLabelsMutex.Lock()
	fake.setLabelsArgsForCall = append(fake.setLabelsArgsForCall, struct {
		arg1 []core.Label
	}{arg1Copy})
	stub := fake.SetLabelsStub
	fake.recordInvocation("SetLabels", []interface{}{arg1Copy})
	fake.setLabelsMutex.Unlock()
	if stub != nil {
		fake.SetLabelsStub(arg1)
	}
}

func (fake *FakePullRequest) SetLabelsCallCount() int {
	fake.setLabelsMutex.RLock()
	defer fake.setLabelsMutex.RUnlock()
	return len(fake.setLabelsArgsForCall)
}

func (fake *FakePullRequest) SetLabelsCalls(stub func([]core.Label)) {
	fake.setLabelsMutex.Lock()
	defer fake.setLabelsMutex.Unlock()
	fake.SetLabelsStub = stub
}

func (fake *FakePullRequest) SetLabelsArgsForCall(i int) []core.Label {
	fake.setLabelsMutex.RLock()
	defer fake.setLabelsMutex.RUnlock()
	argsForCall := fake.setLabelsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePullRequest) SetTargetBranch(arg1 string) {
	fake.setTargetBranchMutex.Lock()
	fake.setTargetBranchArgsForCall = append(fake.setTargetBranchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetTargetBranchStub
	fake.recordInvocation("SetTargetBranch", []interface{}{arg1})
	fake.setTargetBranchMutex.Unlock()
	if stub != nil {
		fake.SetTargetBranchStub(arg1)
	}
}

func (fake *FakePullRequest) SetTargetBranchCallCount() int {
	fake.setTargetBranchMutex.RLock()
	defer fake.setTargetBranchMutex.RUnlock()
	return len(fake.setTargetBranchArgsForCall)
}

func (fake *FakePullRequest) SetTargetBranchCalls(stub func(string)) {
	fake.setTargetBranchMutex.Lock()
	defer fake.setTargetBranchMutex.Unlock()
	fake.SetTargetBranchStub = stub
}

func (fake *FakePullRequest) SetTargetBranchArgsForCall(i int) string {
	fake.setTargetBranchMutex.RLock()
	defer fake.setTargetBranchMutex.RUnlock()
	argsForCall := fake.setTargetBranchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePullRequest) SetTitle(arg1 string) {
	fake.setTitleMutex.Lock()
	fake.setTitleArgsForCall = append(fake.setTitleArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetTitleStub
	fake.recordInvocation("SetTitle", []interface{}{arg1})
	fake.setTitleMutex.Unlock()
	if stub != nil {
		fake.SetTitleStub(arg1)
	}
}

func (fake *FakePullRequest) SetTitleCallCount() int {
	fake.setTitleMutex.RLock()
	defer fake.setTitleMutex.RUnlock()
	return len(fake.setTitleArgsForCall)
}

func (fake *FakePullRequest) SetTitleCalls(stub func(string)) {
	fake.setTitleMutex.Lock()
	defer fake.setTitleMutex.Unlock()
	fake.SetTitleStub = stub
}

func (fake *FakePullRequest) SetTitleArgsForCall(i int) string {
	fake.setTitleMutex.RLock()
	defer fake.setTitleMutex.RUnlock()
	argsForCall := fake.setTitleArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePullRequest) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBodyMutex.RLock()
	defer fake.getBodyMutex.RUnlock()
	fake.getCommitBranchMutex.RLock()
	defer fake.getCommitBranchMutex.RUnlock()
	fake.getLabelsMutex.RLock()
	defer fake.getLabelsMutex.RUnlock()
	fake.getTargetBranchMutex.RLock()
	defer fake.getTargetBranchMutex.RUnlock()
	fake.getTitleMutex.RLock()
	defer fake.getTitleMutex.RUnlock()
	fake.setBodyMutex.RLock()
	defer fake.setBodyMutex.RUnlock()
	fake.setCommitBranchMutex.RLock()
	defer fake.setCommitBranchMutex.RUnlock()
	fake.setLabelsMutex.RLock()
	defer fake.setLabelsMutex.RUnlock()
	fake.setTargetBranchMutex.RLock()
	defer fake.setTargetBranchMutex.RUnlock()
	fake.setTitleMutex.RLock()
	defer fake.setTitleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePullRequest) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ core.PullRequest = new(FakePullRequest)
