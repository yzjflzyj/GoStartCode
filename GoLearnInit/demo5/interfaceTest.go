package main

import (
	"fmt"
)

// WalkRun 接口
type WalkRun interface {
	Walk()
	Run()
}

// Walk
// @Description:
// @receiver p
func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}
func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}
