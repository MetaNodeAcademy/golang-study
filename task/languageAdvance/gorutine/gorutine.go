package main

import (
	"fmt"
	"sync"
	"time"
)

//gorutine
/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
*/

func printJS() {
	var wg sync.WaitGroup
	wg.Add(2) //几个gorutine就启动几个
	go func() {
		defer wg.Done() //当前gorutine执行完毕后调用，同时减1
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("当前偶数为：%d\n", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("当前奇数为：%d\n", i)
			}
		}
	}()
	wg.Wait() //等待所有gorutine执行完毕
}

// 任务结构体
type Task struct {
	Name string
	Job  func()
	Time time.Duration
}

// 调度器结构体
type Scheduler struct {
	tasksQueue chan Task
	workers    int
	wg         sync.WaitGroup
}

// 创建调度器
func newScheduler(workers int, queueSize int) *Scheduler {
	return &Scheduler{
		tasksQueue: make(chan Task, queueSize),
		workers:    workers,
	}
}

// 添加任务到调度器
func (s *Scheduler) addTask(name string, job func()) {
	s.wg.Add(1)
	s.tasksQueue <- Task{Name: name, Job: job}
}

// 启动调度器
func (s *Scheduler) start() {
	for i := 0; i < s.workers; i++ {
		go func() {
			for task := range s.tasksQueue {
				s.execute(task)
			}
		}()
	}
}

// 执行任务
func (s *Scheduler) execute(task Task) {
	defer s.wg.Done()
	start := time.Now()
	fmt.Printf("开始执行任务：%s\n", task.Name)
	task.Job() //执行实际的函数
	fmt.Printf("任务%s执行完毕，耗时：%s\n", task.Name, time.Since(start))
}

// 等待所有任务完成
func (s *Scheduler) wait() {
	close(s.tasksQueue)
	s.wg.Wait()
}
func main() {
	scheduler := newScheduler(3, 10)
	scheduler.start()
	//添加测试任务

	scheduler.addTask("任务1", func() {
		time.Sleep(time.Second * 1)
		// for i := 0; i < 100; i++ {
		// 	fmt.Println("任务1正在执行...")
		// }
		fmt.Println("任务1执行完毕")
	})
	scheduler.addTask("任务2", func() {
		time.Sleep(time.Second * 2)
		// for i := 0; i < 1000; i++ {
		// 	fmt.Println("任务2正在执行...")
		// }
		fmt.Println("任务2执行完毕")
	})
	scheduler.addTask("任务3", func() {
		time.Sleep(time.Second * 3)
		// for i := 0; i < 10000; i++ {
		// 	fmt.Println("任务3正在执行...")
		// }
		fmt.Println("任务3执行完毕")
	})

	//等待所有任务完成
	scheduler.wait()
	fmt.Println("所有任务完成")
}
