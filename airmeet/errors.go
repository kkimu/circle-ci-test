package main

import (
	"fmt"
)

// Error codes
const (
	ErrCodeNotExist      = 1
	ErrCodeAlreadyExists = 2
)

// Error The serializable Error structure.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NewError creates an error instance with the specified code and message.
func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

// Success 成功
type Success struct {
	Result interface{} `json:"result"`
	Code   int         `json:"code"`
}

// NewSuccess 新しいSuccessを作成
func NewSuccess(result interface{}) *Success {
	return &Success{
		Result: result,
		Code:   200,
	}
}

// MajorMessage MajorとMessageを返す
type MajorMessage struct {
	Major   int
	Message string
}
