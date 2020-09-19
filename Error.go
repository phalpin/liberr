package liberr

import (
	"bytes"
	"liberr/errortype"
	"runtime/debug"
)

//#region Error Structs
type Error struct {
	Message    string
	StackTrace string
	ErrorType  errortype.ErrorType
}

type KnownError struct {
	*Error
	FriendlyMessage string
}

//#endregion

//#region Known Error Inits
func NewKnownFromErr(err error, friendlyMsg string, opts ...Option) *KnownError {
	return NewKnown(err.Error(), friendlyMsg, opts...)
}

func NewKnown(msg string, friendlyMsg string, opts ...Option) *KnownError {
	baseErr := New(msg, opts...)

	retVal := &KnownError{
		Error:           baseErr,
		FriendlyMessage: friendlyMsg,
	}

	return retVal
}

//#endregion

//#region Base Error Inits
func New(msg string, opts ...Option) *Error {
	options := &options{}

	if opts != nil {
		for _, val := range opts {
			val(options)
		}
	}

	buf := new(bytes.Buffer)
	buf.Write(debug.Stack())
	retVal := &Error{
		Message:    msg,
		StackTrace: buf.String(),
		ErrorType:  options.errorType,
	}

	return retVal
}

func NewFromError(err error, opts ...Option) *Error {
	return New(err.Error(), opts...)
}

//#endregion