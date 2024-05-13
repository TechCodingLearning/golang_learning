package logSyntax

import (
	"log"
)

// log 包相关联的标志，用来控制写到每个日志项的其他信息
/*
const (
	// Ldate 日期： 2009/01/23
	Ldate = 1 << iota            // 1 << 0

	// Ltime 时间：01:23:23
	Ltime                        // 1 << 1

	// Lmicroseconds 毫秒级时间：01:23:23.123123 该设置会覆盖Ltime标志
	Lmicroseconds                // 1 << 2

	// Llongfile 完整路径的文件名和行号： /a/b/c/d.go:23
	Llongfile                     // 1 << 3

	// Lshortfile 最终的文件名元素和行号：d.go:23 覆盖Llongfile
	Lshortfile                    // 1 << 4

	// LstdFlags 标准日志记录器的初始值
	LstdFlags = Ldate | Ltime
)
*/

// init函数比较特殊，会在程序执行时先执行，同一个包的init顺序
func init() {
	log.SetPrefix("TRACE:") // log日志前缀
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

/*
log 包有一个很方便的地方就是，这些日志记录器是多 goroutine 安全的。
这意味着在多个 goroutine 可以同时调用来自同一个日志记录器的这些函数，而不 会有彼此间的写冲突。
标准日志 记录器具有这一性质，用户定制的日志记录器也应该满足这一性质。
*/

func TestLog() {
	// Println 写到标注内置记录器
	log.Println("message")

	// Fatalln 在调用Println()之后会接着调用os.Exit()
	//log.Fatalln("fatal message")
	defer func() {
		recover() // recover必须在defer中使用才能生效
		log.Println("recover!")
	}()
	// Panicln 在调用Println()之后会接着调用panic()
	log.Panicln("panic message")

}
