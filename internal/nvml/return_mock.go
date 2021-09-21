// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package nvml

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"sync"
)

// Ensure, that ReturnMock does implement Return.
// If this is not the case, regenerate this file with moq.
var _ Return = &ReturnMock{}

// ReturnMock is a mock implementation of Return.
//
// 	func TestSomethingThatUsesReturn(t *testing.T) {
//
// 		// make and configure a mocked Return
// 		mockedReturn := &ReturnMock{
// 			ErrorFunc: func() string {
// 				panic("mock out the Error method")
// 			},
// 			StringFunc: func() string {
// 				panic("mock out the String method")
// 			},
// 			ValueFunc: func() nvml.Return {
// 				panic("mock out the Value method")
// 			},
// 		}
//
// 		// use mockedReturn in code that requires Return
// 		// and then make assertions.
//
// 	}
type ReturnMock struct {
	// ErrorFunc mocks the Error method.
	ErrorFunc func() string

	// StringFunc mocks the String method.
	StringFunc func() string

	// ValueFunc mocks the Value method.
	ValueFunc func() nvml.Return

	// calls tracks calls to the methods.
	calls struct {
		// Error holds details about calls to the Error method.
		Error []struct {
		}
		// String holds details about calls to the String method.
		String []struct {
		}
		// Value holds details about calls to the Value method.
		Value []struct {
		}
	}
	lockError  sync.RWMutex
	lockString sync.RWMutex
	lockValue  sync.RWMutex
}

// Error calls ErrorFunc.
func (mock *ReturnMock) Error() string {
	callInfo := struct {
	}{}
	mock.lockError.Lock()
	mock.calls.Error = append(mock.calls.Error, callInfo)
	mock.lockError.Unlock()
	if mock.ErrorFunc == nil {
		var (
			sOut string
		)
		return sOut
	}
	return mock.ErrorFunc()
}

// ErrorCalls gets all the calls that were made to Error.
// Check the length with:
//     len(mockedReturn.ErrorCalls())
func (mock *ReturnMock) ErrorCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockError.RLock()
	calls = mock.calls.Error
	mock.lockError.RUnlock()
	return calls
}

// String calls StringFunc.
func (mock *ReturnMock) String() string {
	callInfo := struct {
	}{}
	mock.lockString.Lock()
	mock.calls.String = append(mock.calls.String, callInfo)
	mock.lockString.Unlock()
	if mock.StringFunc == nil {
		var (
			sOut string
		)
		return sOut
	}
	return mock.StringFunc()
}

// StringCalls gets all the calls that were made to String.
// Check the length with:
//     len(mockedReturn.StringCalls())
func (mock *ReturnMock) StringCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockString.RLock()
	calls = mock.calls.String
	mock.lockString.RUnlock()
	return calls
}

// Value calls ValueFunc.
func (mock *ReturnMock) Value() nvml.Return {
	callInfo := struct {
	}{}
	mock.lockValue.Lock()
	mock.calls.Value = append(mock.calls.Value, callInfo)
	mock.lockValue.Unlock()
	if mock.ValueFunc == nil {
		var (
			returnOut nvml.Return
		)
		return returnOut
	}
	return mock.ValueFunc()
}

// ValueCalls gets all the calls that were made to Value.
// Check the length with:
//     len(mockedReturn.ValueCalls())
func (mock *ReturnMock) ValueCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockValue.RLock()
	calls = mock.calls.Value
	mock.lockValue.RUnlock()
	return calls
}
