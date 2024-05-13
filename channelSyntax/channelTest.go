package channelSyntax

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 关闭信道
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			// 注意看这个输出，可以发现还是存在阻塞的，因为当c没有输出，缓冲没空，且quit没有输入，就会出现这种状态！
			fmt.Println("waiting...")
		}
	}
}

/*
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

ch := make(chan int)
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。
*/

func TestChannelNormal() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // 普通信道
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从c接收

	fmt.Println(x, y, x+y)

}

func TestChannelCache() {
	/*
		带缓冲的信道
		信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：

		ch := make(chan int, 100)
		仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
	*/
	ch := make(chan int, 2) // 带缓冲的信道
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func TestChanelClose() {
	/*
		发送者可通过 close 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

		v, ok := <-ch
		之后 ok 会被设置为 false。

		循环 for i := range c 会不断从信道接收值，直到它被关闭。

		*注意：* 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

		*还要注意：* 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
	*/
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
select 语句使一个 Go 程可以等待多个通信操作。

select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
*/

func TestChannelSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

//type Tree struct {
//	Left  *Tree
//	Value int
//	Right *Tree
//}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)

}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	isPass := true
	for i := 0; i < 10; i++ { // 让Walk函数能够走完，否则没读完会阻塞到main函数结束，浪费资源
		if x, y := <-ch1, <-ch2; x != y {
			isPass = false
		}
	}
	return isPass
}

func TestChannelPracticeCase() {
	// 测试Walk步进tree
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// 测试Same检测t1和t2是否含有相同的值
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
