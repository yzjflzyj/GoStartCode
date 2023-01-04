package main

import "fmt"

func main() {

	//1.指针类型不能声明后再赋值，因为指针类型没有默认分配内存。而值类型会默认分配内存，并给一个零值。
	//指针类型可以new，主动分配一块内存，并返回这块内存的指针
	var sp *string
	sp = new(string) //关键点
	*sp = "飞雪无情"
	fmt.Println(*sp)

	//2.make 函数不只是 map 类型的工厂函数，还是 chan、slice 的工厂函数。它同时可以用于 slice、chan 和 map 这三种类型的初始化。

	//new和make的区别：new只是主动分配内存，并返回内存指针，没有做初始化。make在new的基础上，还做了初始化。
}
