package main

import (
	"fmt"
	"sync"
	"time"
)

//共享的资源
var sum = 0

func main() {
	// sum为非同步安全变量
	//syncError()

	//利用sync.Mutex的锁保证安全
	//syncSafe()

	//读锁的使用
	//syncRead()

	// sync.WaitGroup的使用，类似countLatch
	//run()

	//sync.Once：只执行一次的场景
	//doOnce()

	//sync.NewCond:类似cyclicbarrier
	race()
}

//不安全写
func syncError() {
	//开启500个协程让sum+10
	for i := 0; i < 500; i++ {
		go add(10)
	}
	//防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为:", sum)
}

func add(i int) {
	sum += i
}

//安全写
func syncSafe() {
	//开启500个协程让sum+10
	for i := 0; i < 500; i++ {
		go addForSafe(10)
	}
	//防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为:", sum)
}

var mutex sync.Mutex

func addForSafe(i int) {
	mutex.Lock()
	//确保不会遗忘
	defer mutex.Unlock()
	sum += i
}

//读写锁，在保证不读到脏数组的同时，提升性能，多个读锁可以同时读取
func syncRead() {
	for i := 0; i < 500; i++ {
		go addForSafe(10)
	}
	for i := 0; i < 10; i++ {
		go fmt.Println("和为:", readSum())
	}
	time.Sleep(2 * time.Second)
	fmt.Println("和为:", sum)
}

var mutexRw sync.RWMutex

// 读锁的使用
func readSum() int {
	//只获取读锁
	mutexRw.RLock()
	defer mutexRw.RUnlock()
	b := sum
	return b
}

// sync.WaitGroup的使用，类似countLatch
func run() {
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(510)
	for i := 0; i < 500; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			addForSafe(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}
	//一直等待，只要计数器值为0
	wg.Wait()
}

//sync.Once：只执行一次的场景
func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	//用于等待协程执行完毕
	done := make(chan bool)
	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	fmt.Println(<-done)
}

//sync.NewCond，有如下三个方法，它的三个方法 Wait、Signal、Broadcast 就分别对应 Java 中的 wait、notify、notifyAll。
//1. Wait，阻塞当前协程，直到被其他协程调用 Broadcast 或者 Signal 方法唤醒，使用的时候需要加锁，使用 sync.Cond 中的锁即可，也就是 L 字段。
//2. Signal，唤醒一个等待时间最长的协程。
//3. Broadcast，唤醒所有等待的协程。
//10个人赛跑，1个裁判发号施令
func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	var wgReady sync.WaitGroup
	wgReady.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			wgReady.Done() //确保都进入等待
			cond.Wait()    //等待发令枪响
			fmt.Println(num, "号开始跑……")
			cond.L.Unlock()
		}(i)
	}

	//等待所有goroutine都进入wait状态
	wgReady.Wait()
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() //发令枪响
	}()

	//防止函数提前返回退出
	wg.Wait()
}
