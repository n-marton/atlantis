// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/runatlantis/atlantis/server (interfaces: GitlabRequestParser)

package mocks

import (
	http "net/http"
	"reflect"

	pegomock "github.com/petergtz/pegomock"
)

type MockGitlabRequestParser struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGitlabRequestParser() *MockGitlabRequestParser {
	return &MockGitlabRequestParser{fail: pegomock.GlobalFailHandler}
}

func (mock *MockGitlabRequestParser) Validate(r *http.Request, secret []byte) (interface{}, error) {
	params := []pegomock.Param{r, secret}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Validate", params, []reflect.Type{reflect.TypeOf((*interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockGitlabRequestParser) VerifyWasCalledOnce() *VerifierGitlabRequestParser {
	return &VerifierGitlabRequestParser{mock, pegomock.Times(1), nil}
}

func (mock *MockGitlabRequestParser) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierGitlabRequestParser {
	return &VerifierGitlabRequestParser{mock, invocationCountMatcher, nil}
}

func (mock *MockGitlabRequestParser) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierGitlabRequestParser {
	return &VerifierGitlabRequestParser{mock, invocationCountMatcher, inOrderContext}
}

type VerifierGitlabRequestParser struct {
	mock                   *MockGitlabRequestParser
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierGitlabRequestParser) Validate(r *http.Request, secret []byte) *GitlabRequestParser_Validate_OngoingVerification {
	params := []pegomock.Param{r, secret}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Validate", params)
	return &GitlabRequestParser_Validate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type GitlabRequestParser_Validate_OngoingVerification struct {
	mock              *MockGitlabRequestParser
	methodInvocations []pegomock.MethodInvocation
}

func (c *GitlabRequestParser_Validate_OngoingVerification) GetCapturedArguments() (*http.Request, []byte) {
	r, secret := c.GetAllCapturedArguments()
	return r[len(r)-1], secret[len(secret)-1]
}

func (c *GitlabRequestParser_Validate_OngoingVerification) GetAllCapturedArguments() (_param0 []*http.Request, _param1 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*http.Request, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*http.Request)
		}
		_param1 = make([][]byte, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]byte)
		}
	}
	return
}
