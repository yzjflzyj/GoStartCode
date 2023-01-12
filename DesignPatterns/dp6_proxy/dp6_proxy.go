package main

import "fmt"

//抽象主题

type Subject interface {
	Request() string
}

// 真实主题

type RealSubject struct{}

func (RealSubject) Request() string {
	fmt.Print("real")
	return "real"
}

// 代理

type Proxy struct {
	real RealSubject
}

func NewProxy(subject RealSubject) *Proxy {
	return &Proxy{
		real: subject,
	}
}

func (Proxy) Pre() {
	fmt.Print("pre:")
}

func (Proxy) After() {
	fmt.Print(":after")
}

func (p Proxy) Request() string {
	// 代理的额外处理
	p.Pre()
	// 调用真实对象
	p.real.Request()
	// 调用之后的操作
	p.After()
	return "proxy"
}

func main() {
	realSubject := RealSubject{}
	proxy := NewProxy(realSubject)
	proxy.Request()
}
