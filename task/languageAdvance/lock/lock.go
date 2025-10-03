package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
// func main() {
// 	var lock sync.Mutex
// 	var wg sync.WaitGroup
// 	var totalNum int
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(num int) {
// 			defer wg.Done()
// 			for j := 0; j < 1000; j++ {
// 				lock.Lock()
// 				totalNum++
// 				lock.Unlock()
// 				fmt.Printf("gorutine %d is running\n", num)
// 			}
// 		}(i)
// 	}
// 	wg.Wait()
// 	fmt.Printf("totalNum: %d\n", totalNum)

// }

var counter int64
var wg sync.WaitGroup

func main() {
	var WorkerNum = 10 //10个gorutine
	var LoopNum = 1000 //每个gorutine执行1000次

	wg.Add(WorkerNum)
	for i := 0; i < WorkerNum; i++ {
		go func(id int) {
			fmt.Printf("gorutine %d is running\n", id)
			defer wg.Done()
			for j := 0; j < LoopNum; j++ {
				atomic.AddInt64(&counter, 1) //原子递增操作
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("counter: %d\n", counter)

}
