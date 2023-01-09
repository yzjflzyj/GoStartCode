package main

import (
	"fmt"
	"sync"
)

// get，set方法忽略
type singleton struct {
	Value int
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 利用sync.once实现单例模式
func GetInstance(v int) *singleton {
	once.Do(func() {
		instance = &singleton{Value: v}
	})
	return instance
}
func main() {
	s1 := GetInstance(1)
	s2 := GetInstance(2)
	if s1 != s2 {
		fmt.Println("实例不相同")
	}
	fmt.Println(s1.Value, s2.Value)
}
