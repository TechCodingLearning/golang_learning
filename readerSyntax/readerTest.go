package readerSyntax

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
	"os"
	"strings"
)

/*
io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。

Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。

io.Reader 接口有一个 Read 方法：

func (T) Read(b []byte) (n int, err error)
Read 用数据 "填充给定的字节切片" 并 "返回填充的字节数" 和 "错误值" 。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
*/

type MyReader struct {
}

func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func TestReader() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	reader.Validate(MyReader{})

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	mr := rot13Reader{s}
	io.Copy(os.Stdout, &mr)
}

type rot13Reader struct {
	r io.Reader
}

// 转换byte 前进13位/后退13位
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b += 13
	case 'M' < b && b <= 'Z':
		b -= 13
	case 'a' <= b && b <= 'm':
		b += 13
	case 'm' < b && b <= 'z':
		b -= 13
	}
	return b
}

// 重写Read方法
func (mr rot13Reader) Read(b []byte) (int, error) {
	n, e := mr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}
