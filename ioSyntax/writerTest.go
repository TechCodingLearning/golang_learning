package ioSyntax

import (
	"bytes"
	"fmt"
	"os"
)

func TestWriter() {
	// 创建一个Buffer值，并将一个字符串写入Buffer
	// 使用实现io.Writer的Write方法
	var b bytes.Buffer
	b.Write([]byte("Hello"))

	// 使用Fprintf来将一个字符串拼接到Buffer里
	// 将bytes.Buffer的地址作为io.Writer类型值传入
	fmt.Fprintf(&b, "World!")

	// 将Buffer的内容输出到标准输出设备
	// 将io.File值的地址作为io.Writer类型值传入
	b.WriteTo(os.Stdout)
}
