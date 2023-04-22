package main

import (
	"errors"
	"fmt"
)

func main() {
	sum := sum(1, 3)
	fmt.Println(sum)

	//匿名函数
	sum2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum2(1, 2))

	//函数闭包,得到的匿名函数,可以访问匿名函数外的函数的变量
	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	//调用方法
	age := Age(25)
	age.String()

	//指针类型接受者
	age.Modify()
	age.String()

	//方法表达式调用
	//方法赋值给变量，方法表达式
	sm := Age.String
	//通过变量，要传一个接收者进行调用也就是age
	sm(age)

}

// 函数声明
func sum(a, b int) int {
	return a + b
}

// 多值返回
func sumAndError(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	return a + b, nil
}

// 返回参数命名,不常用
func sumForNameResult(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

// 可变参数,实际就是切片,有普通参数时,可变参数要放在最后
func sum1(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

// 返回匿名函数
func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Age uint //表示定义一个新类型 Age，该类型等价于 uint

// 方法声明,方法 String() 就是类型 Age 的方法,类型 Age 是方法 String() 的接收者。
func (age Age) String() {
	fmt.Println("the age is", age)
}

// Modify 指针类型接受者
func (age *Age) Modify() {
	*age = Age(30)
}
