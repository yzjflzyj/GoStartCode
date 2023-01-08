package main

import (
	"fmt"
)

//定义接口
type WalkRun interface {
	Walk()
	Run()
}

//实现接口的方法，要实现全部方法
func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}
func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}
