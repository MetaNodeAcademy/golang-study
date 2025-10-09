package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
		编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
		并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
	*/

	var wg sync.WaitGroup
	wg.Add(2)
	channel := make(chan int)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel) //这里须关闭channel，否则会死锁
	}()

	go func() {
		defer wg.Done()
		for v := range channel {
			fmt.Printf("接收到数据:%d\n", v)
		}
	}()
	wg.Wait()

	/*
		题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
	*/
	channel2 := make(chan int, 10)
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go func() {
		defer wg2.Done()
		for i := 0; i < 100; i++ {
			channel2 <- i
		}
		close(channel2)
	}()

	go func() {
		defer wg2.Done()
		for value := range channel2 {
			fmt.Printf("接收到数据:%d\n", value)
		}
	}()
	wg2.Wait()
	/////////////////////////////////////////////////
	const (
		totalTaskNum = 100
		bufferSize   = 10
		consumerNum  = 2
	)
	ch := make(chan int, bufferSize)
	var wg3 sync.WaitGroup
	wg3.Add(1)
	//启动生产者
	go producer(ch, totalTaskNum, &wg3)
	//启动多个消费者
	for i := 0; i < consumerNum; i++ {
		wg3.Add(1)
		go consumer(ch, i, &wg3)
	}
	wg3.Wait()
	fmt.Println("所有任务处理完成")

}

// 生产者
func producer(producerChan chan<- int, totalTaskNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < totalTaskNum; i++ {
		producerChan <- i
		fmt.Printf("生产者生产了数据:%d\n", i)
		time.Sleep(time.Millisecond * 200)
	}
	close(producerChan)
}

// 消费者
func consumer(consumerChan <-chan int, consumerNo int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range consumerChan {
		fmt.Printf("消费者%d正在消费数据:%d\n", consumerNo, value)
		time.Sleep(time.Millisecond * 400)
	}
}
