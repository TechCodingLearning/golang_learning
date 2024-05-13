package errorSyntax

import (
	"fmt"
	"math"
	"time"
)

/*
error类型是一个内建接口：
type error interface{
	Error() string
}
*/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error { // 返回一个error接口
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	const eps = 1e-6

	var z float64 = 1.0
	var z_before float64 = z

	for math.Abs(z-z_before) > eps {
		z_before = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
	//return 0, nil
}

func TestError() {
	if err := run(); err != nil { // err为nil表示成功，反之为失败
		fmt.Println(err) // 内部调用err.Error()方法
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
