package main

import "fmt"

// Component 抽象构件
type Component interface {
	Calc() int
}

// ConcreteComponent 具体构件
type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

// MulDecorator 装饰者角色1
type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

// AddDecorator 装饰者角色2
type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

func main() {
	concreteComponent := &ConcreteComponent{}
	mulDecorator := WrapMulDecorator(concreteComponent, 1)
	fmt.Println(mulDecorator.Calc())
	addDecorator := WrapAddDecorator(concreteComponent, 2)
	fmt.Println(addDecorator.Calc())
}
