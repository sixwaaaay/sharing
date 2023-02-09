// Package errorx
// 用于本项目中的错误处理，
// 错误码和错误信息可以用于响应给客户端的错误信息
package errorx

const defaultCode = 1001

type CodeError struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"`
}

func (e *CodeError) Error() string {
	return *e.StatusMsg
}

func NewCodeError(code int64, msg string) error {
	return &CodeError{StatusCode: code, StatusMsg: &msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func NewDefaultCodeErr(msg string) *CodeError {
	return &CodeError{StatusCode: defaultCode, StatusMsg: &msg}
}

func (e *CodeError) Data() *CodeError {
	return &CodeError{StatusCode: e.StatusCode, StatusMsg: e.StatusMsg}
}
