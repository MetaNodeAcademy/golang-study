package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {

	var wg sync.WaitGroup //锁，用于同步gorutine

	wg.Add(3) //这里的数字表示子gorutine数量,即子gorutine的计数器,有点类似java中的countDownLatch

	go func() {
		defer wg.Done() //此处类似java中的countDownLatch的countDown操作
		fmt.Println("gorutine runnig...")
	}()

	go func(ss string) {
		defer wg.Done()
		fmt.Println(ss)
	}("gorutine-1 running...")

	go func() {
		defer wg.Done()
		say("gorutine loop...")
	}()

	// go say("gorutine loop...") //这里在没有引入同步机制的情况下可能只打印一次，也就是子gorutine还没执行完，但主gorutine已执行完

	say("hello")
	wg.Wait() //阻塞等待所有子gorutine执行完毕 ，有点类似于java中的countDownLatch.await()

	////////////////////////////////////////////////////////////////////////////////////////////////

	counter := SafeConter{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.increment()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("SafeConter获取到的结果为：%d\n", counter.count)

	counter1 := UnSafeConter{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter1.increment()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("SafeConter获取到的结果为：%d\n", counter1.count)

}

// 线程安全的计数器
type SafeConter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeConter) increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 减少计数
func (c *SafeConter) decrement() {
	c.mu.Lock()
	c.count--
	c.mu.Unlock()
}

// 获取当前计数
func (c *SafeConter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 非线程安全的计数器
type UnSafeConter struct {
	count int
}

func (c *UnSafeConter) increment() {
	c.count++
}
func (c *UnSafeConter) decrement() {
	c.count--
}
func (c *UnSafeConter) getValue() int {
	return c.count
}
