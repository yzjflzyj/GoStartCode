package main

import "fmt"

// Builder 建造者接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Director 指挥者
type Director struct {
	builder Builder
}

// NewDirector 构造函数：构造指挥者
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 指挥者执行建造
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// ConcreteBuilder 具体建造者
type ConcreteBuilder struct{}

func (b *ConcreteBuilder) Part1() {
	fmt.Println("part1")
}

func (b *ConcreteBuilder) Part2() {
	fmt.Println("part2")
}

func (b *ConcreteBuilder) Part3() {
	fmt.Println("part3")
}

func main() {
	builder := &ConcreteBuilder{}
	director := NewDirector(builder)
	director.Construct()
}
