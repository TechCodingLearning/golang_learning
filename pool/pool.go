/*
@author: louris
@since: 2022/8/1
@desc: //TODO 使用带缓冲的通道实现资源池
*/

package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理了一组可以安全地在多个goroutine间共享的资源
// 被管理的资源必须实现io.Closer接口
type Pool struct {
	mux       sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 表示请求(Acquire)了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

// New
// @Description //TODO New 创建一个用来管理资源的池
// @Param fn func() (io.Closer, error)：可以分配新资源的函数
// @Param size uint：规定池的大小
// @return
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和Close操作的安全
	p.mux.Lock()
	defer p.mux.Unlock()

	// 如果池已经被关闭，销毁这个资源
	if p.closed {
		err := r.Close()
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	select {
	// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		err := r.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// Close 会让资源停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	// 保证本操作与Release操作的安全
	p.mux.Lock()
	defer p.mux.Unlock()

	// 如果pool已经被关闭，什么也不做
	if p.closed {
		return
	}

	// 将池关闭
	p.closed = true

	// 在清空通道里的资源之前，将通道关闭
	// 如果不这样做，会发生死锁
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		err := r.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
