package errorwrapper

import "fmt"

// E for creating new error.
// error should always be the first param.
func E(args ...interface{}) error {
	if len(args) == 0 {
		return Errorf("errorx.E: bad call without args")
	}

	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case *Error:
			// Copy and put the errors back.
			errCopy := *arg
			e = &errCopy

		case error:
			e.Err = arg

		case string:
			e.Err = Errorf(arg)

		case Code:
			// New code will always replace the old code.
			e.Code = arg

		case Message:
			e.Message = arg

		default:
			// The default error is unknown.
			msg := fmt.Sprintf("errorwrapper.E: bad call, args=%v", args)
			return Errorf(msg+"; unknown_type=%T value=%v", arg, arg)
		}
	}
	return e
}

func (m Message) String() string {
	return string(m)
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	if e.Err == nil {
		return "bad call without args"
	}

	return e.Err.Error()
}

// Error returns the string representation of the error message.
func (e *errorString) Error() string {
	return e.s
}

// Errorf is equivalent to fmt.Errorf, but allows clients to import only this
// package for all error handling.
func Errorf(format string, args ...interface{}) error {
	return &errorString{
		s: fmt.Sprintf(format, args...),
	}
}
