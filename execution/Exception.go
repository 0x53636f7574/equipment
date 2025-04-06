package execution

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type StackTrace = string

type ExceptionHandler func(exception *Exception)

type Exception struct {
	Reason error
	Trace  StackTrace
}

func (exception *Exception) String() string {
	return fmt.Sprintf("Error: %v.\n%v", exception.Reason.Error(), exception.Trace)
}

func (exception *Exception) Error() string {
	return exception.Reason.Error()
}

func Revert(reason error) {
	panic(&Exception{Reason: reason})
}

func Trace(reason error) {
	panic(&Exception{Reason: reason, Trace: string(debug.Stack())})
}

func Catch(handler ExceptionHandler) {
	if err := recover(); err != nil {
		if e, ok := err.(error); ok {
			var exc *Exception
			if errors.As(e, &exc) {
				handler(exc)
			} else {
				handler(&Exception{Reason: e})
			}
		}
	}
}
