package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1.变量定义:var 变量名 类型 = 表达式
	//定义的变量必须使用,否则编译不通过
	var i int = 10
	//类型推导功能
	var j = 10
	//多个变量同时声明
	var (
		k int = 1
		m int = 2
	)
	var (
		n = 0
		o = 1
	)
	//变量简写形式
	s1 := "Hello"
	fmt.Println("Hello, 世界!")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k + m)
	fmt.Println(n + o)
	fmt.Println(s1)
	//常量
	const name = "飞雪无情"
	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)
	//int和string互转
	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)
	//float和string进行互转
	f1, _ := strconv.ParseFloat("1222.3456", 2)
	f2 := strconv.FormatFloat(0.125, 'f', 2, 64)
	fmt.Println(f1, f2)
	//bool和string互转
	b1, _ := strconv.ParseBool("true")
	b2 := strconv.FormatBool(false)
	fmt.Println(b1, b2)
	//类型转
	i2f := float64(i)
	f2i := int(f1)
	fmt.Println(i2f, f2i)

	//strings包
	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1, "H"))
	fmt.Println(strings.HasSuffix(s1, "o"))
	//在s1中查找字符串o
	fmt.Println(strings.Index(s1, "o"))
	//把s1全部转为大写
	fmt.Println(strings.ToUpper(s1))
}
