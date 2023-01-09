package main

import "fmt"

// 利用值类型的参数传递是拷贝，实现原型模式
type prototype struct {
	id   string
	name string
}

func (p *prototype) Clone() *prototype {
	//获得值类型，这个值是一个拷贝
	value := *p
	//返回指针，就是在内存中分配一块
	return &value
}

func main() {
	t1 := &prototype{
		id:   "1",
		name: "type1",
	}

	t2 := t1.Clone()

	if t1 == t2 {
		fmt.Println("原型模式出错")
	}
}
