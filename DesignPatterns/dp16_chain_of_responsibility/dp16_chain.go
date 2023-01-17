package main

import "fmt"

// Manager 处理者接口
type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

// RequestChain 责任链：处理者容器，也是处理者接口的实现者，控制流转
type RequestChain struct {
	Manager
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (r *RequestChain) HaveRight(money int) bool {
	return true
}

// ProjectManager 具体处理者1
type ProjectManager struct{}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (*ProjectManager) HaveRight(money int) bool {
	return money > 1000
}

func (*ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "刘备" {
		fmt.Printf("Project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request\n", name, money)
	return false
}

// DepManager 具体处理者2
type DepManager struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{},
	}
}

func (*DepManager) HaveRight(money int) bool {
	return money < 1000 && money > 100
}

func (*DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "关羽" {
		fmt.Printf("Dep manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request\n", name, money)
	return false
}

// GeneralManager 具体处理者3
type GeneralManager struct{}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{},
	}
}

func (*GeneralManager) HaveRight(money int) bool {
	return true
}

func (*GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "张飞" {
		fmt.Printf("General manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request\n", name, money)
	return false
}

func main() {
	projectManagerChain := NewProjectManagerChain()
	depManagerChain := NewDepManagerChain()
	generalManagerChain := NewGeneralManagerChain()

	projectManagerChain.SetSuccessor(depManagerChain)
	depManagerChain.SetSuccessor(generalManagerChain)

	var c Manager = projectManagerChain
	c.HandleFeeRequest("刘备", 1001)
	c.HandleFeeRequest("关羽", 800)
	c.HandleFeeRequest("张飞", 100)
}
