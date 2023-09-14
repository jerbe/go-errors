package errors

import (
	"errors"
	"fmt"
	"testing"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/8/20 23:05
  @describe :
*/

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
