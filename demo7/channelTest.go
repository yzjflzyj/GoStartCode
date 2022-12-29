package main

import (
	"fmt"
	"time"
)

func main() {
	// 协程Goroutine的使用
	go fmt.Println("飞雪无情")
	fmt.Println("我是 main goroutine")
	time.Sleep(time.Second)

	// 协程的通信：channel
	//1. 接收：获取 chan 中的值，操作符为 <- chan。
	//2. 发送：向 chan 发送值，把值放在 chan 中，操作符为 chan <-。
	channelTest()

	//select+channel
	selectChannel()
}

// channel：协程通信
func channelTest() {
	// 无缓冲channel，没有容量，会阻塞
	ch := make(chan string)
	// 有缓冲channel
	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))

	//关闭channel：如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。
	//但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。
	close(cacheCh)

	// 单向channel:若进行非定义的操作，会编译报错
	onlySend := make(chan<- int)
	onlyReceive := make(<-chan int)
	fmt.Println(onlySend, onlyReceive)

	go func() {
		fmt.Println("飞雪无情")
		// 发送
		ch <- "goroutine 完成"
	}()

	fmt.Println("我是 channelTest goroutine")

	// 接收
	v := <-ch
	fmt.Println("channelTest：接收到的chan中的值为：", v)
}

func selectChannel() {
	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)
	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		threeCh <- downloadFile("threeCh")
	}()
	//开始select多路复用，哪个channel能获取到值，
	//就说明哪个最先下载好，就用哪个。
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}
func downloadFile(chanName string) string {
	//模拟下载文件,可以自己随机time.Sleep点时间试试
	time.Sleep(time.Second)
	return chanName + ":filePath"
}
