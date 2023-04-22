package main

import "fmt"

func main() {
	//结构体声明,和其他变量一样
	p := person{"飞雪无情", 30, address{"广东", "深圳"}}
	p1 := person{age: 30, name: "飞雪无情"}
	//获取字段值
	fmt.Println(p, p.name, p.age, p1)

	//接口
	//函数调用,函数入参是接口,p实现了接口
	printString(p)
	//值类型的接受者实现接口,其指针也自动实现; 但是反过来,指针类型的接受者实现了接口,其值类型的并没有自动实现
	printString(&p)

	//工厂函数
	p2 := NewPerson("张三")
	fmt.Println(p2)

	//组合结构体的声明
	p3 := person1{
		age:  30,
		name: "飞雪无情",
		address: address{
			province: "北京",
			city:     "北京",
		},
	}
	//像使用自己的字段一样，直接使用
	fmt.Println(p3.province)
}

// 结构体定义
type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

// 接口的实现
func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

// 函数,参数为接口, person实现了Stringer接口,该函数是面向接口编程的
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// NewPerson 工厂函数,创建一个指针类型的结构体
func NewPerson(name string) *person {
	return &person{name: name}
}

// 结构体的组合(接口也可以组合)
type person1 struct {
	name string
	age  uint
	address
}
