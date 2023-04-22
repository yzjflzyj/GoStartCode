package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
*
转换规则：
1. 任何类型的 *T 都可以转换为 unsafe.Pointer；
2. unsafe.Pointer 也可以转换为任何类型的 *T；
3. unsafe.Pointer 可以转换为 uintptr；
4. uintptr 也可以转换为 unsafe.Pointer。
unsafe.Pointer 主要用于指针类型的转换，而且是各个指针类型转换的桥梁。uintptr 主要用于指针运算，尤其是通过偏移量定位不同的内存。
*/
func main() {
	//unsafe.Pointer
	reflectTest()
	//两个指针是不能转换的，强制转换，需要使用 unsafe 包里的 Pointer
	unsafePointerTest()
	//uintptr：地址偏移
	uintptrTest()
	//unsafe.Size：大小只与类型有关，和类型对应的变量存储的内容大小无关，比如 bool 型占用一个字节、int8 也占用一个字节。
	sizeOfTest()
}

// 一个 struct 结构体的内存占用大小，等于它包含的字段类型内存占用大小之和。
func sizeOfTest() {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
	fmt.Println(unsafe.Sizeof(string("飞雪无情")))
	fmt.Println(unsafe.Sizeof([]string{"飞雪u无情", "张三"}))
}

// 通过地址偏移，找到内存，进行操作
func uintptrTest() {
	p := new(person)
	//Name是person的第一个字段不用偏移，即可通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "飞雪无情"
	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)
}

func unsafePointerTest() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
}

func reflectTest() {
	p := person{Name: "飞雪无情", Age: 20}
	pv := reflect.ValueOf(p)
	//反射调用person的Print方法
	mPrint := pv.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)
}

func (p person) Print(prefix string) {
	fmt.Printf("%s:Name is %s,Age is %d\n", prefix, p.Name, p.Age)
}

type person struct {
	Name string
	Age  int
}
