package gcpmetadata

import "fmt"

type Error interface {
	Code() ErrCode
	error
}

type appError struct {
	C ErrCode
	M string
}

func (e *appError) Code() ErrCode {
	return e.C
}

func (e *appError) Error() string {
	return e.M
}

type ErrCode int

const (
	ErrUnknownCode ErrCode = iota
	ErrNotFoundCode
	ErrInvalidArgumentCode
)

func ErrNotFound(msg string) Error {
	return &appError{C: ErrNotFoundCode, M: msg}
}

func ErrInvalidArgument(expected string, argument string) Error {
	return &appError{C: ErrInvalidArgumentCode, M: fmt.Sprintf("invalid argument. expected is %v, argument is = %v", expected, argument)}
}

func Is(err error, code ErrCode) bool {
	e, ok := err.(Error)
	if !ok {
		return false
	}
	return e.Code() == code
}
