package liberr

import "github.com/phalpin/liberr/errortype"

type options struct {
	errorType errortype.ErrorType
}

type Option = func(opts *options)

func WithErrorType(errType errortype.ErrorType) Option {
	return func(opts *options) {
		opts.errorType = errType
	}
}
