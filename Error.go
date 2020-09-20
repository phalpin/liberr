package liberr

import (
	"bytes"
	"github.com/phalpin/liberr/errortype"
	"runtime/debug"
)

//#region Error Structs

type BaseError struct {
	Message    string
	StackTrace string
	ErrorType  errortype.ErrorType
}

type KnownError struct {
	*BaseError
	FriendlyMessage string
}

//#endregion

//#region Required Implementation for Error
func (err *BaseError) Error() string {
	return err.Message
}

func (err *KnownError) Error() string {
	return err.Message
}

//#endregion

//#region Known Error Inits
func NewKnownFromErr(err error, friendlyMsg string, opts ...Option) *KnownError {
	return NewKnown(err.Error(), friendlyMsg, opts...)
}

func NewKnown(msg string, friendlyMsg string, opts ...Option) *KnownError {
	baseErr := NewBase(msg, opts...)

	retVal := &KnownError{
		BaseError:       baseErr,
		FriendlyMessage: friendlyMsg,
	}

	return retVal
}

//#endregion

//#region Base Error Inits
func NewBase(msg string, opts ...Option) *BaseError {
	options := &options{}

	if opts != nil {
		for _, val := range opts {
			val(options)
		}
	}

	buf := new(bytes.Buffer)
	buf.Write(debug.Stack())
	retVal := &BaseError{
		Message:    msg,
		StackTrace: buf.String(),
		ErrorType:  options.errorType,
	}

	return retVal
}

func NewBaseFromError(err error, opts ...Option) *BaseError {
	return NewBase(err.Error(), opts...)
}

//#endregion
