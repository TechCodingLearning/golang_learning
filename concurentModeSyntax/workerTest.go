package concurentModeSyntax

import (
	"GolangProjects/work"
	"log"
	"sync"
	"time"
)

// names 提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func TestWorker() {
	// 使用两个goroutine来创建工作池
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 迭代names切片
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}

			go func() {
				// 将任务提交执行
				// 当Run返回时，我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	p.Shutdown()
}
