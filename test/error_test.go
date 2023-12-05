package test

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestError(t *testing.T) {
	a := 10
	b := 0
	c, d := div(a, b)
	if err, ok := d.(error); ok {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
func div(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("除数为0")
		return 0, err
	} else {
		return a / b, nil
	}
}
