package main

import "fmt"

func main() {
	point()
}

func point() {
	name := "关羽"
	nameP := &name //取地址
	fmt.Println("name变量的值为:", name)
	fmt.Println("name变量的内存地址为:", nameP)

	//1.指针的声明
	/**
	var intP *int
	intP = &name //此时intP是nil类型，指针类型不同，无法赋值
	*/
	//因此用var声明而不赋值的方式都不可以
	p := 10
	//完整声明：var intP *int = new(int)
	intP1 := new(int) //简短声明
	intP1 = &p
	fmt.Println(intP1)

	//2.指针操作，会修改对应的值
	*nameP = "张飞"
	fmt.Println(name)

	//3.函数传值，都是对值的拷贝，无法在另一个函数中对入参进行修改。若是指针类型，则可以对数据进行修改，因为传的是内存的地址的拷贝，对应的内存不变。

	//4.使用指针的好处：
	//1. 可以修改指向数据的值；
	//2. 在变量赋值，参数传值的时候可以节省内存。
}
