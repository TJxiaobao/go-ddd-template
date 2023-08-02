package errno

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Error struct {
	code        int
	message     string
	args        []interface{}
	originError error
	caller      string
	callStack   string
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Message() string {
	if len(e.args) == 0 {
		return e.message
	}
	return fmt.Sprintf(e.message, e.args...)
}

func (e Error) Error() string {
	return fmt.Sprintf("{\"OriginError\": \"%v\", \"Message\": \"%s\", \"Caller\": \"%s\", \"CallStack\": \"%s\"}", e.originError, e.Message(), e.caller, e.callStack)
}

func (e Error) IsSuccess() bool {
	return e.code == OK.Code
}

func NewError(no *Errno, err error, args ...interface{}) error {
	if e, ok := err.(Error); ok {
		return e
	}
	if e, ok := err.(*Error); ok {
		return *e
	}
	return Error{
		code:        no.Code,
		message:     no.Message,
		args:        args,
		originError: err,
		caller:      getCaller(),
		callStack:   getStack(),
	}
}

func NewSimpleError(no *Errno, err error, args ...interface{}) error {
	if e, ok := err.(Error); ok {
		return e
	}
	if e, ok := err.(*Error); ok {
		return *e
	}
	return Error{
		code:        no.Code,
		message:     no.Message,
		args:        args,
		originError: err,
	}
}

func AssertError(err error) Error {
	if err == nil {
		return Error{
			code:    OK.Code,
			message: OK.Message,
		}
	}
	if e, ok := err.(Error); ok {
		return e
	}
	if e, ok := err.(*Error); ok {
		return *e
	}
	return Error{
		originError: err,
		code:        ErrUnknown.Code,
		message:     ErrUnknown.Message,
	}
}

func getStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	splitStr := "Error"
	sa := strings.SplitN(string(buf[:n]), splitStr, 2)
	if len(sa) == 2 {
		subStr := sa[1]
		splitStr := "\n"
		sa2 := strings.SplitN(subStr, splitStr, 3)
		if len(sa2) == 3 {
			s := "\n" + sa2[2]
			// Fixme splitStr需要修改为middleware函数名
			splitStr := "main.ApiWrap.func1"
			sa3 := strings.SplitN(s, splitStr, 2)
			if len(sa3) == 2 {
				return sa3[0]
			}
		}
	}

	return string(buf[:n])
}

func getCaller() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)

	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	return fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)
}
