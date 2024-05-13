/*
@author: louris
@since: 2022/8/1
@desc: //TODO 使用通道来监视程序的执行时间
*/
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete 通道报告处理任务已经超时
	complete chan error // 正常完成返回nil，发生错误返回err

	// timeout 报告处理任务已经超时
	timeout <-chan time.Time // 单向，只接收

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d), // d时间后返回一条time.Time的通道消息
	}
}

// Add 将一个任务附加到Runner上，这个任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
		// 如果没有default，select将阻塞，直到某个case可以执行
		//default:
		//	fmt.Println("default\n")
		//	return nil
	}

}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
