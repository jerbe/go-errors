# GO-Errors 一个简单易用的golang错误包
## 说明

本包是可以用来追踪错误链,能够打印出错误的包裹链,对于收集日志,bug调试特别有用

## 使用方法
```go
import (
    "errors"
    "fmt"
    "testing"
)


var T = New("this is test")

func TestWrap(t *testing.T) {
    var e error
    e = errors.New("Error One")
    e = T.Wrap(e)
    e = e.(*Error).Wrap(e)
    e = Wrap(e)
    e = Wrap(e)
    e = Wrap(e)
    fmt.Printf("%+v\n", e)
    fmt.Printf("%s\n", e.Error())
    
    e2 := New("Error Two")
    e = e.(*Error).Wrap(e2)
    e = Wrap(e)
    e = Wrap(e)
    fmt.Printf("%+v", e)
    fmt.Printf("%s\n", e.Error())
}


//OUTPUT
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:24
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:23
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:22
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:21 => this is test
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:20 => this is test
Error One
Error One
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:31
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:30
[github.com/jerbe/go-errors.TestWrap]:/Users/Jerbe/Workshop/Study/github.com/jerbe/go-errors/error_test.go:29
Error Two
Error Two
```