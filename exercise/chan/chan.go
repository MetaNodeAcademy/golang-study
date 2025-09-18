package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("开始学习channle")

	str := "我是局部变量"

	func(num int) int {
		fmt.Println(str)
		fmt.Println(num + 2)
		return num + 1
	}(2) //这里的匿名函数的()表示立即执行，也可以传入匿名函数的实参

	time.Sleep(1 * time.Second)

	fmt.Println("执行到这里了")

	//不带缓冲的通道
	channel := make(chan string, 1) //这里必须设置缓冲，否则会阻塞
	channel <- "hello sam"
	val, ok := <-channel
	if ok {
		fmt.Printf("读取到channel中的值为：%v\n", val)
	} else {
		fmt.Println("channel中没有数据")
	}

	//不设置缓冲的话可以这样操作
	channel2 := make(chan string)
	go func() {
		//写数据
		channel2 <- "hello jack"

	}()
	val2, success := <-channel2
	if success {
		fmt.Printf("读取到channel2中的值为：%v\n", val2)
	} else {
		fmt.Println("channel2中没有数据")
	}

	var channel5 = make(chan int, 4)
	//写入数据
	channel5 <- 888
	go func() {
		num66 := <-channel5
		fmt.Printf("从channel5中读取到的数据为：%v\n", num66)
	}()

	//只读通道
	channel3 := make(<-chan string)
	//channel3 <-"hello world"  这里会报错
	go func() {
		val, ok := <-channel3
		if ok {
			fmt.Printf("从channel3中读取到的数据为：%v\n", val)
		} else {
			fmt.Println("channel3中没有数据")
		}
	}()

	// //只写通道
	// channel4:=make(chan<- string)

	dataChan := producer()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		consumer(dataChan)
	}()

	wg.Wait() // 等待消费者goroutine完成
	fmt.Println("所有数据处理完成")

	chan666 := make(chan string, 8)
	go producer1(chan666)
	go consumer1(chan666)

	//主线程睡眠一下等待协程执行
	time.Sleep(2 * time.Second)

}

// 数据生产者
func producer() <-chan string {
	channel := make(chan string)
	go func() {
		defer close(channel)
		for i := 0; i < 5; i++ {
			channel <- fmt.Sprintf("生产者生产了第%d个数据", i)
		}
	}()
	return channel
}

// 数据消费者
func consumer(channel <-chan string) {
	for data := range channel {
		fmt.Println("消费者消费了数据:", data)
	}
}

func producer1(ch chan string) {
	fmt.Println("生产者启动")
	ch <- "这是第一个数据"
	ch <- "这是第二个数据"
	ch <- "这是第三个数据"
	fmt.Println("生产者结束")
}

func consumer1(ch chan string) {
	fmt.Println("消费者启动")
	for {
		data := <-ch
		fmt.Println("消费者消费了数据:", data)
	}
}
