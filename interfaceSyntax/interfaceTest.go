package interfaceSyntax

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func Test() {

	//_interface.TestInterface()
	//var a _interface.Abser
	//f := _interface.MyFloat(-math.Sqrt2)
	//v := _interface.Vertex{3, 4}
	//fmt.Println(v.Abs())
	//a = f  // a MyFloat 实现了 Abser
	//a = &v // a *Vertex 实现了 Abser
	//
	//// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	//// 所以没有实现 Abser。
	//a = &v
	//
	//fmt.Println(a.Abs())
	//
	//t := _interface.T{"hello"}
	//t.M()

	// 接口也是值
	var a Abser

	var v *Vertex
	a = v
	Describe(a)
	fmt.Println(a.Abs())
	a = &Vertex{3, 4}
	Describe(a)
	fmt.Println(a.Abs())

	a = MyFloat(math.Pi)
	Describe(a)
	fmt.Println(a.Abs())

	var i interface{}
	DescribeAnyType(i)

	i = 42
	DescribeAnyType(i)

	i = "hello"
	DescribeAnyType(i)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	if v == nil { // 即便接口内的具体值为nil，方法仍然会被nil接收者调用，go通常会写一些方法来优雅地粗粝空指针nil问题
		fmt.Println("<nil>")
		return -1
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestInterface() {
	fmt.Println("test interfaceSyntax...")
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() { //此方法表示类型T实现了接口I，但我们无需显示声明此事。
	fmt.Println(t.S)
}

func Describe(a Abser) { // 大写才能被外包使用
	fmt.Printf("(%v, %T)", a, a)
}

func DescribeAnyType(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
