package routineSyntax

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

go f(x, y, z)
会启动一个新的 Go 程并执行

f(x, y, z)
f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中。

Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestRoutine() {
	go say("world")
	say("hello")
}

func TestRoutineNoShift() {
	// 分配一个逻辑处理器给调度器使用
	// 1 运行太快，没来得及切换routine以至于顺序输出，
	// 2 可以清楚地看到并发
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Printf("\n")
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Printf("\n")
	}()

	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")

}

var wg sync.WaitGroup

func TestRoutineShift() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Wating to finish")
	wg.Wait()
	fmt.Println("Terminating program")
}

func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				time.Sleep(1000)
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
