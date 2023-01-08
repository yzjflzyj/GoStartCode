package main

import "fmt"

func main() {
	add := address{province: "北京", city: "北京"}
	//1.值实现了接口，那么对应指针类型也实现了接口，但是反过来不成立
	printString(add)
	printString(&add)
	//2.值所实现的接口，可以当入参；值实现的接口的指针，不能当入参，因为指向接口的指针没有实现接口（所以永远不要使用接口的指针，没有用处）
	var si fmt.Stringer = address{province: "上海", city: "上海"}
	printString(si)
	sip := &si
	//printString(sip)报错
	fmt.Println(sip)

	//3.引用类型：map、slice、channel，函数，接口 其本质是指针类型，属于语法糖，因此可以传值修改，go本身没有引用类型。在此只是一种习惯性称呼
	//make是go的内建函数，最终调用的都是 runtime.makemap 函数。
	m := make(map[string]int)
	m["飞雪无情"] = 18
	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])

	//4.值类型的零值：如图，没有显式初始化，则默认零值
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

type address struct {
	province string
	city     string
}

//实现了fmt.Stringer接口
func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}
