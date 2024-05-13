package baseSyntax

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"
)

//var c, python, java bool
//
//var (
//	ToBe   bool       = false
//	MaxInt uint64     = 1<<64 - 1
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//)

const Pi = 3.14 // 不能用:=语法声明

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //直接返回返回值定义的值
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g\n", v, lim)
	}
	// 不可到达
	return lim
}

func Sqrt(x float64) float64 {
	const eps = 1e-6

	var z float64 = 1.0
	var z_before float64 = z

	for math.Abs(z-z_before) > eps {
		z_before = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Vertex1 struct {
	X, Y float64
}

func (v Vertex1) Abs() float64 { // 给结构体定义方法
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Test() {
	//fmt.Println("My favorite number is", rand.Intn(10))
	//fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//fmt.Println(math.Sqrt(6))
	//fmt.Println(math.Pi)
	//fmt.Println(add(42, 13))
	//
	//a, b := swap("hello", "world")
	//fmt.Println(a, b)
	//
	//fmt.Println(split(17))
	//
	//var i, j int = 1, 2
	//k := 3
	//c, python, java := true, false, "no!"
	//fmt.Println(i, j, k, c, python, java)
	//
	//fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	//fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	//fmt.Printf("Type: %T Value: %v\n", z, z)

	//var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)
	//fmt.Println(x, y, z)
	//
	//var j = x
	//
	//fmt.Println(j)
	//fmt.Printf("Type: %T\n", j)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// for{} // 无限循环
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20))

	fmt.Println(Sqrt(2))

	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// fallthrough 默认带有break，想要连续执行则使用fallthrough
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch { // 相当于 switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("world") // 将函数推迟到外层函数返回之后执行，推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用

	fmt.Println("hello")

	fmt.Println("counting")

	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i) // 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
	//}
	//fmt.Println("done")
	//
	//i, j := 42, 2701
	//fmt.Printf("Type: %T\n", i)
	//p := &i
	//fmt.Println(*p)
	//
	//*p = 21
	//fmt.Println(i)
	//
	//p = &j
	//*p = *p / 37 // go中没有指针运算，例如p++
	//fmt.Println(j)
	//
	//fmt.Println(Vertex{1, 2})
	//
	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v.X)
	//fmt.Println(v)
	//
	//ptr := &v
	//ptr.X = 1e9
	//fmt.Println(v)
	//
	//var pp *Vertex
	//pp = &v
	//fmt.Println(*pp)
	//
	//var (
	//	v1 = Vertex{1, 2}
	//	v2 = Vertex{X: 1}
	//	v3 = Vertex{}
	//	px = &Vertex{1, 2}
	//)
	//fmt.Println(v1, px, v2, v3)

	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)

	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)

	//var s []int = primes[1:4] //切片
	//fmt.Println(s)

	//// 切片并不存储任何数据，类似于引用
	//names := [4]string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//fmt.Println(names)
	//
	//a := names[0:2]
	//b := names[1:3]
	//fmt.Println(a, b)
	//
	//b[0] = "XXX"
	//fmt.Println(a, b)
	//fmt.Println(names)
	//
	//q := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println(q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)
	//
	//s := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, true},
	//	{7, true},
	//	{11, false},
	//	{13, true},
	//}
	//fmt.Println(s)
	//
	//q = q[:2]
	//printSlice(q)
	//fmt.Println(q[:cap(q)])
	//
	//var x []int
	//printSlice(x)
	//if x == nil {
	//	fmt.Println("nil!")
	//}
	//
	//var y = []int{}
	//printSlice(y)
	//if y == nil {
	//	fmt.Println("y nil!")
	//}

	// make创建动态数组
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)

	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)

	var xs = []int{3, 2, 1, 5, 6}

	for i, v := range xs {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	//var m map[string]Vertex
	//m = make(map[string]Vertex)
	//m["a"] = Vertex{1, 2}
	//fmt.Println(m["a"])
	//
	//mm := map[string]Vertex{
	//	"a": {1, 2},
	//	"b": {2, 3},
	//}
	//fmt.Println(mm)

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) //key不存在，类型默认值

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i))
	}

	v1 := Vertex1{3, 4}
	fmt.Println(v1.Abs())
}
