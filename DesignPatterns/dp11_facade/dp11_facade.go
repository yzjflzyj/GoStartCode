package main

import "fmt"

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API is facade interface 总接口
type API interface {
	Test() string
}

// apiImpl facade implement 实现的结构体及其实现方法
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// 接口A及其实现结构体，实现方法，构造函数

// NewAModuleAPI return new AModuleAPI 构造函数
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI ... 接口
type AModuleAPI interface {
	TestA() string
}

// 结构体
type aModuleImpl struct{}

// TestA 实现方法
func (*aModuleImpl) TestA() string {
	return "A module running"
}

// 接口B及其实现结构体，实现方法，构造函数

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

func main() {
	api := NewAPI()
	fmt.Println(api.Test())
}
