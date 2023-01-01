package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
* Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。
* 一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。主要有四个方法：
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
1. Deadline 方法可以获取设置的截止时间，第一个返回值 deadline 是截止时间，到了这个时间点，Context 会自动发起取消请求，第二个返回值 ok 代表是否设置了截止时间。
2. Done 方法返回一个只读的 channel，类型为 struct{}。在协程中，如果该方法返回的 chan 可以读取，则意味着 Context 已经发起了取消信号。通过 Done 方法收到这个信号后，就可以做清理操作，然后退出协程，释放资源。
3. Err 方法返回取消的错误原因，即因为什么原因 Context 被取消。
4. Value 方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 key 才可以获取对应的值。

四种context：
WithCancel(parent Context)：生成一个可取消的 Context。
WithDeadline(parent Context, d time.Time)：生成一个可定时取消的 Context，参数 d 为定时取消的具体时间。
WithTimeout(parent Context, timeout time.Duration)：生成一个可超时取消的 Context，参数 timeout 用于设置多久后取消
WithValue(parent Context, key, val interface{})：生成一个可携带 key-value 键值对的 Context。
*/

/**
Context 使用原则
Context 是一种非常好的工具，使用它可以很方便地控制取消多个协程。在 Go 语言标准库中也使用了它们，比如 net/http 中使用 Context 取消网络的请求。
要更好地使用 Context，有一些使用原则需要尽可能地遵守。
1. Context 不要放在结构体中，要以参数的方式传递。
2. Context 作为函数的参数时，要放在第一位，也就是第一个参数。
3. 要使用 context.Background 函数生成根节点的 Context，也就是最顶层的 Context。
4. Context 传值要传递必须的值，而且要尽可能地少，不要什么都传。
5. Context 多协程安全，可以在多个协程中放心使用。
以上原则是规范类的，Go 语言的编译器并不会做这些检查，要靠自己遵守。
*/
func main() {
	//select+channel的方式实现监控直到完成
	//watchUntilFinish()

	//context+select的方式，传递停止信息
	//contextAndSelect()

	//context传值
	contextWithValue()
}

//完成时退出监控
func watchUntilFinish() {
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan bool) //用来停止监控狗
	go func() {
		defer wg.Done()
		watchDog(stopCh, "【监控狗1】")
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stopCh <- true              //发停止指令
	wg.Wait()
}

//用select+channel完成通知
func watchDog(stopCh chan bool, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-stopCh:
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}

//context+select的方式，传递停止信息
func contextAndSelect() {
	var wg sync.WaitGroup
	wg.Add(3)
	// 生成的是可取消的context，自身context及其子context都会收到取消信号
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDogForContext(ctx, "【监控狗1】")
	}()
	go func() {
		defer wg.Done()
		watchDogForContext(ctx, "【监控狗2】")
	}()
	go func() {
		defer wg.Done()
		watchDogForContext(ctx, "【监控狗3】")
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令
	wg.Wait()
}

func watchDogForContext(ctx context.Context, name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}

//传值型子context
func contextWithValue() {
	var wg sync.WaitGroup
	wg.Add(4) //记得这里要改为4，原来是3，因为要多启动一个协程
	//省略其他无关代码
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDogForContext(ctx, "【监控狗1】")
	}()
	go func() {
		defer wg.Done()
		watchDogForContext(ctx, "【监控狗2】")
	}()
	go func() {
		defer wg.Done()
		watchDogForContext(ctx, "【监控狗3】")
	}()
	//增加传值型子context
	valCtx := context.WithValue(ctx, "userId", 2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()
	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令
	wg.Wait()
	//省略其他无关代码
}

func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】", "协程退出")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("【获取用户】", "用户ID为：", userId)
			time.Sleep(1 * time.Second)
		}
	}
}
