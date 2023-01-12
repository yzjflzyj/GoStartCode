package main

import "fmt"

// SaveArticle 抽象模式工厂接口，抽象工厂
type SaveArticle interface {
	CreateProse() Prose
	CreateAncientPoetry() AncientPoetry
}

// 具体工厂1
type RedisFactory struct{}

// 生产抽象产品Prose
func (*RedisFactory) CreateProse() Prose {
	return &RedisProduct{}
}

// 生产抽象产品AncientPoetry
func (*RedisFactory) CreateAncientPoetry() AncientPoetry {
	return &RedisProduct{}
}

// 具体工厂2
type MysqlFactory struct{}

// 生产抽象产品Prose
func (*MysqlFactory) CreateProse() Prose {
	return &MysqlProduct{}
}

// 生产抽象产品AncientPoetry
func (*MysqlFactory) CreateAncientPoetry() AncientPoetry {
	return &MysqlProduct{}
}

// 抽象产品1：Prose 散文
type Prose interface {
	SaveProse()
}

// 抽象产品2：AncientPoetry 古诗
type AncientPoetry interface {
	SaveAncientPoetry()
}

// 具体产品1
type RedisProduct struct{}

func (*RedisProduct) SaveProse() {
	fmt.Println("RedisProduct Save Prose")
}

func (*RedisProduct) SaveAncientPoetry() {
	fmt.Println("RedisProduct Save Ancient Poetry")
}

// 具体产品2
type MysqlProduct struct{}

func (*MysqlProduct) SaveProse() {
	fmt.Println("MysqlProduct Save Prose")
}

func (*MysqlProduct) SaveAncientPoetry() {
	fmt.Println("MysqlProduct Save Ancient Poetry")
}

// 工厂方法
func Save(saveArticle SaveArticle) {
	saveArticle.CreateProse().SaveProse()
	saveArticle.CreateAncientPoetry().SaveAncientPoetry()
}
func main() {
	var factory SaveArticle
	factory = &RedisFactory{}
	Save(factory)
	factory = &MysqlFactory{}
	Save(factory)
}
