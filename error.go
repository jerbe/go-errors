package go_errors

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/8/20 11:46
  @describe :
*/

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"sync/atomic"
)

var errID = uint64(0)

// newErrorID 获取错误ID
func newErrorID() uint64 {
	return atomic.AddUint64(&errID, 1)
}

// Error 栈错误
type Error struct {
	cause   error
	message string
	caller  string
	id      uint64
}

// Is 判断target是否是同一个,比如e被克隆出d来就能用Is来判断两个是否相等
func (e *Error) Is(target error) bool {
	if t, ok := target.(*Error); ok {
		return t.id == e.id && t.message == e.message
	}
	return false
}

// Error 错误的文本
func (e *Error) Error() string {
	if e.cause != nil {
		msg := e.message
		if msg == "" {
			return e.cause.Error()
		}
		msg += " <= " + e.cause.Error()
		return msg
	}
	return e.message
}

// Cause 上一个error,用到Wrap时有用
func (e *Error) Cause() error {
	return e.cause
}

// Wrap 可以用来包裹其他错误
func (e *Error) Wrap(target error) error {
	ne := *e
	ne.cause = target
	ne.caller = caller()
	return &ne
}

// Unwrap 解包错误
func (e *Error) Unwrap() error {
	return e.cause
}

func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		// 打印出具体的调用栈,当使用Wrap时才有用
		// 用法 fmt.Sprintf("%+v",err) 等用格式化输出的都可以
		if s.Flag('+') {
			msg := ""
			if e.message != "" && e.caller != "" {
				msg = " => "
			}
			msg += e.message
			fmt.Fprintf(s, "%s%s\n", e.caller, msg)

			if e.cause != nil {
				fmt.Fprintf(s, "%+v", e.cause)
			}

			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())

	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// New 生成错误
func New(message string) *Error {
	return &Error{message: message, id: newErrorID()}
}

// NewWithCaller 生成带caller的错误
func NewWithCaller(message string) *Error {
	return &Error{message: message, id: newErrorID(), caller: caller()}
}

func caller() string {
	pc, file, line, _ := runtime.Caller(2)
	return "[" + runtime.FuncForPC(pc).Name() + "]:" + file + ":" + strconv.Itoa(line)
}
