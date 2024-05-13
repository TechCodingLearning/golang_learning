package typeSyntax

import "fmt"

// TestTypeAssertion 类型断言
// t := i.(T)
// 如果i保存了具体类型T，则返回底层类型为T的值；反之，触发panic
// t, ok := i.(T)
// 如果i保存了具体类型T，则返回具体的值和true；反之，返回类型T的零值和false，不会触发panic
func TestTypeAssertion() {

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) //
	fmt.Println(f, ok)

	//f = i.(float64) // 产生panic
	//fmt.Println(f)
}

// TestTypeSelection 类型选择
// 类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。
// 此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。
func TestTypeSelection(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

// TestStringer fmt包中定义的接口
// Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
func TestStringer() {
	a := Person{"Arthur Dent", 32}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z) // 直接调用String()

	var i fmt.Stringer
	i = a
	fmt.Println(i.String())
	i = z
	fmt.Println(i.String())
}
