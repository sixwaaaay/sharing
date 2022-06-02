package errorx

import (
	"github.com/jinzhu/copier"
	"testing"
)

type resp struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"`
}

func TestNewCodeError(t *testing.T) {
	testCode := int64(200)
	testMsg := "test"
	err := NewCodeError(testCode, testMsg)
	codeError := err.(*CodeError)
	var res resp
	copier.Copy(&res, codeError)
	if res.StatusCode != testCode {
		t.Errorf("code error status code is not %d", testCode)
	}
	if res.StatusMsg == nil || *res.StatusMsg != testMsg {
		t.Errorf("code error status msg is not %s", testMsg)
	}
}

func TestNewDefaultCodeErr(t *testing.T) {
	testMsg := "test"
	codeError := NewDefaultCodeErr(testMsg)
	var res resp
	copier.Copy(&res, codeError)
	if res.StatusCode != defaultCode {
		t.Errorf("default code error status code is not %d", defaultCode)
	}
	if res.StatusMsg == nil || *res.StatusMsg != "test" {
		t.Errorf("default code error status msg is not %s", testMsg)
	}
}

func TestNewDefaultError(t *testing.T) {
	testMsg := "msg"
	err := NewDefaultError(testMsg)
	codeError := err.(*CodeError)
	var res resp
	copier.Copy(&res, codeError)
	if res.StatusCode != defaultCode {
		t.Errorf("default code error status code is not %d", defaultCode)
	}
	if res.StatusMsg == nil || *res.StatusMsg != "msg" {
		t.Errorf("default code error status msg is not %s", testMsg)
	}
}
