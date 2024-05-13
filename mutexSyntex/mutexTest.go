package mutexSyntex

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock() // Lock之后同一时刻只有一个goroutine能访问c.v
	return c.v[key]
}

func TestMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey" + string(i))
	}

	//time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
