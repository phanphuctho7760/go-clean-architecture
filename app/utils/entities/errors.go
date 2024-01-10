package entities

import (
	"fmt"
	"runtime/debug"
)

type ErrorCustom struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func WrapError(err error, messagef string, msgArgs ...interface{}) ErrorCustom {
	return ErrorCustom{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}
func (err ErrorCustom) Error() string { return err.Message }
