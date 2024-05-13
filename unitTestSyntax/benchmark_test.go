package unitTestSyntax

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf 对fmt.Sprintf进行基准测试
func BenchmarkSprintf(b *testing.B) {
	number := 10

	// 开始循环之前需要进行初始化，这个方法用来重置计时器，
	// 保证测试代码执行前的初始化代码不会干扰计时器的结果
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat 对strconv.FormatInt函数进行基准测试
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 对strconv.Itoa函数进行基准测试
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
