package liberr

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_BasicCases(t *testing.T) {
	testMsg := "this is just a test"
	err := New(testMsg)
	assert.Equal(t, Unknown, err.ErrorType)
	assert.Equal(t, testMsg, err.Message)
	assert.NotEqual(t, "", err.StackTrace)

	testMsg = "this is another test......."
	err = New(testMsg)
	assert.Equal(t, Unknown, err.ErrorType)
	assert.Equal(t, testMsg, err.Message)
	assert.NotEqual(t, "", err.StackTrace)

	testMsg = "this is a test with a specified error type option"
	err = New(testMsg, WithErrorType(InvalidArgument))
	assert.Equal(t, InvalidArgument, err.ErrorType)
	assert.Equal(t, testMsg, err.Message)
	assert.NotEqual(t, "", err.StackTrace)
}

func Test_NewFromError_BasicCases(t *testing.T) {
	testMsg := errors.New("this is just a test")
	err := NewFromError(testMsg)
	assert.Equal(t, Unknown, err.ErrorType)
	assert.Equal(t, testMsg.Error(), err.Message)
	assert.NotEqual(t, "", err.StackTrace)

	testMsg = errors.New("this is another test.......")
	err = NewFromError(testMsg)
	assert.Equal(t, Unknown, err.ErrorType)
	assert.Equal(t, testMsg.Error(), err.Message)
	assert.NotEqual(t, "", err.StackTrace)

	testMsg = errors.New("this is a test with a specified error type option")
	err = NewFromError(testMsg, WithErrorType(InvalidArgument))
	assert.Equal(t, InvalidArgument, err.ErrorType)
	assert.Equal(t, testMsg.Error(), err.Message)
	assert.NotEqual(t, "", err.StackTrace)
}

func Test_NewKnown_BasicCases(t *testing.T) {
	asserts := func(expMsg string, expFrMsg string, expErrType ErrorType, opts ...Option) {
		err := NewKnown(expMsg, expFrMsg, opts...)
		assert.Equal(t, expErrType, err.ErrorType)
		assert.Equal(t, expMsg, err.Message)
		assert.Equal(t, expFrMsg, err.FriendlyMessage)
		assert.NotEqual(t, "", err.StackTrace)
	}

	asserts("this is just a test", "this is a friendly error msg", Unknown)
	asserts("this is another test......", "this is another friendly error msg....", Unknown)
	asserts(
		"this is a test with an error type specified",
		"please do some stuff with this error type",
		InvalidArgument,
		WithErrorType(InvalidArgument),
	)
	asserts(
		"something went horribly wrong",
		"a team of highly trained flying monkeys has been dispatched to deal with this problem",
		Unknown,
		WithErrorType(Unknown),
	)
}

func Test_NewKnownFromErr_BasicCases(t *testing.T) {
	asserts := func(expErr error, expFrMsg string, expErrType ErrorType, opts ...Option) {
		err := NewKnownFromErr(expErr, expFrMsg, opts...)
		assert.Equal(t, expErrType, err.ErrorType)
		assert.Equal(t, expErr.Error(), err.Message)
		assert.Equal(t, expFrMsg, err.FriendlyMessage)
		assert.NotEqual(t, "", err.StackTrace)
	}

	asserts(
		errors.New("stuff happened oh no"),
		"something went down please try again",
		Unknown,
	)

	asserts(
		errors.New("some other stuff happened oh no"),
		"something else went down. please try again",
		Unknown,
	)

	asserts(
		errors.New("invalid user input"),
		"Your information was incomplete. Please enter the proper values and try again.",
		InvalidArgument,
		WithErrorType(InvalidArgument),
	)

	asserts(
		errors.New("unknown err with friendly msg"),
		"An unknown error occurred. This issue has been recorded and will be reacted to before the heat death of the universe.",
		Unknown,
		WithErrorType(Unknown),
	)
}
