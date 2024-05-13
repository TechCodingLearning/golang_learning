package concurentModeSyntax

import (
	"GolangProjects/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

/**
 * @Description
 * 程序可以在分配的时间内完成工作，正常种植；
 * 程序没有及时完成工作，自动终止；
 * 接收到操作系统发送的中断事件，程序立刻试图清理状态并停止工作
 **/
func TestRunner() {
	log.Println("Starting work.")

	// w为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
