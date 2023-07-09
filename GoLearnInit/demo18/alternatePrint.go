package main

import (
	"fmt"
	"sync"
)

func printer(wg *sync.WaitGroup, ch chan int, done chan bool, start, end int) {
	defer wg.Done()

	for i := start; i <= end; i += 2 {
		<-ch // 等待通道可用
		fmt.Println(i)
		ch <- 0 // 发送一个值到通道，以唤醒另一个协程
	}

	// 发送完成信号给另一个协程
	done <- true
}

func main() {
	var wg sync.WaitGroup   // 用来关闭ch
	ch := make(chan int, 1) // 带有缓冲区的通道
	done := make(chan bool) // 用来阻塞主线程

	wg.Add(2)

	go printer(&wg, ch, done, 1, 100) // 第一个协程打印奇数
	go printer(&wg, ch, done, 2, 100) // 第二个协程打印偶数

	// 启动第一个协程
	ch <- 0

	// 等待两个协程完成
	go func() {
		wg.Wait()
		close(ch) // 关闭通道，表示结束
	}()

	// 等待完成信号
	<-done
	<-done

	fmt.Println("打印完成")
}
