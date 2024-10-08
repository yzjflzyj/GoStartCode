package main

import "fmt"

//性能优化:
//Go 语言有两部分内存空间：栈内存和堆内存。
//● 栈内存由编译器自动分配和释放，开发者无法控制。
//栈内存一般存储函数中的局部变量、参数等，函数创建的时候，这些内存会被自动创建；函数返回的时候，这些内存会被自动释放。
//● 堆内存的生命周期比栈内存要长，如果函数返回的值还会在其他地方使用，那么这个值就会被编译器自动分配到堆上。
//堆内存相比栈内存来说，不能自动被编译器释放，只能通过垃圾回收器才能释放，所以栈内存效率会很高。

//判断一个变量应该分配到堆上还是栈上,这就需要逃逸分析

// 优化技巧：指针虽然可以减少内存的拷贝，但它同样会引起逃逸
// 第 1 个需要介绍的技巧是尽可能避免逃逸，因为栈内存效率更高，还不用 GC。比如小对象的传参，array 要比 slice 效果好。
// 如果避免不了逃逸，还是在堆上分配了内存，那么对于频繁的内存申请操作，我们要学会重用内存，比如使用 sync.Pool，这是第 2 个技巧。
// 第 3 个技巧就是选用合适的算法，达到高性能的目的，比如空间换时间。
func main() {
	//go循环变量问题：案例1，取地址符
	items := []Item{Item(1), Item(2), Item(3)}
	var all []*Item
	var all2 []Item
	// 在每次存入变量 all 的都是相同的 item， 因此需要在处理前 item := item
	for _, item := range items {
		all2 = append(all2, item)
		all = append(all, &item)
	}

	for _, e := range all {
		e.String()
	}

	for _, e := range all2 {
		e.String()
	}

	//go循环变量问题：案例2，闭包函数
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		prints = append(prints, func() { fmt.Println(v) })
	}
	for _, print := range prints {
		print()
	}
}

type Item uint

func (item Item) String() {
	fmt.Println("the item is", item)
}
