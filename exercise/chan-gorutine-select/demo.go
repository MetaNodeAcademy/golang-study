package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
主程序启动 ​​3 个 worker goroutine​​，每个 worker 从任务通道接收任务并处理。
任务通道为 ​​有缓冲通道​​，容量为 5。
主程序生成 ​​10 个任务​​（整数 ID），发送到任务通道。
使用 select实现以下逻辑：
每个 worker 处理完任务后，向结果通道发送 任务ID。
主程序统计已处理的任务数，若 ​​8 秒内未完成所有任务​​，触发超时并退出。
输出所有已处理的任务 ID，或提示超时。
*/

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		//模拟任务处理时间0-1秒
		time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))
		//将任务id发送到结果通道
		results <- task
		fmt.Printf("worker %d finished task %d\n", id, task)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	//创建任务通道(缓冲区5)和结果通道
	tasks := make(chan int, 5)
	results := make(chan int)
	//同步
	var wg sync.WaitGroup
	//启动3个worker
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	//使用gorutine生成10个任务
	go func() {
		for i := 0; i < 10; i++ {
			tasks <- i
			fmt.Printf("main goroutine send task %d\n", i)
		}
		close(tasks)
	}()

	//select 监听结果和超时
	timeout := time.After(1 * time.Second)
	flag := 0
	for flag < 10 {
		select {
		case taskId := <-results:
			flag++
			fmt.Printf("main goroutine received task %d\n", taskId)
		case <-timeout:
			fmt.Println("timeout")
			return
		}

	}
	fmt.Println("all task finished")
}
