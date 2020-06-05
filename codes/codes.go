package codes

import "fmt"

const (
	FrameWorkError = 1
	BusinessError  = 2
)

const (
	OK             = 0
	ClientMsgError = 100
)

type Error struct {
	Type    int
	Code    int
	Message string
}

func NewFrameWorkError(code int, msg string) *Error {
	return &Error{Type: FrameWorkError, Code: code, Message: msg}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.Type == FrameWorkError {
		return fmt.Sprintf("type: framework, code: %d, message: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("type: business, code: %d, message: %s", e.Code, e.Message)
}
