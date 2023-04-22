package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//数组：需要指定类型和长度，用[len]或者[...]来指定长度
	a1 := [1]string{"关羽"}
	a2 := [2]string{"张飞"}
	fmt.Println(a1, a2)

	//切片，1.全量声明（不要长度），2.借数组声明，3.用make声明
	ss := []string{"飞雪无情", "张三"}
	//append会自动扩容：扩容机制是新建数组，复制，返回新的变量
	ss = append(ss, "李四", "王五")
	fmt.Println(ss)
	//多个切片共用一个底层数组虽然可以减少内存占用，但是如果有一个切片修改内部的元素，其他切片也会受影响。所以在切片作为参数在函数间传递的时候要小心，尽可能不要修改原切片内的元素。

	//切片的本质是 SliceHeader，又因为函数的参数是值传递，所以传递的是 SliceHeader 的副本，而不是底层数组的副本。
	//这时候切片的优势就体现出来了，因为 SliceHeader 的副本内存占用非常少，即使是一个非常大的切片（底层数组有很多元素），也顶多占用 24 个字节的内存，这就解决了大数组在传参时内存浪费的问题。

	//切片的高效还体现在 for range 循环中，因为循环得到的临时变量也是个值拷贝，所以在遍历大的数组时，切片的效率更高。

	//数组和切片在函数传值时的不同
	sliceAndArr()

	//Go 语言通过先分配一个内存再复制内容的方式，实现 string 和 []byte 之间的强制转换。
	stringsAndSliceByte()
	//强制转换是采用重新分配内存的方式，如果字符串过大，内存开销很大
	compulsionTransfer()
	//string 和 []byte 的互转就是一个很好的利用 SliceHeader 结构体的示例，通过它可以实现零拷贝的类型转换，提升了效率，避免了内存浪费。
	zeroCopy1()
	zeroCopy2()
}

func zeroCopy1() {
	//s4 和 s3 的内容是一样的。不一样的是 s4 没有申请新内存（零拷贝），它和变量 b 使用的是同一块内存，
	//因为它们的底层 Data 字段值相同，这样就节约了内存，也达到了 []byte 转 string 的目的。
	s := "飞雪无情"
	b := []byte(s)
	//s3:=string(b)
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Println(s4)
}

func zeroCopy2() {
	////b1 和 b 的内容是一样的，不一样的是 b1 没有申请新内存，而是和变量 s 使用同一块内存，因为它们底层的 Data 字段相同，所以也节约了内存。
	//通过 unsafe.Pointer 把 string 转为 []byte 后，不能对 []byte 修改，
	//比如不可以进行 b1[0]=12 这种操作，会报异常，导致程序崩溃。这是因为在 Go 语言中 string 内存是只读的。
	s := "飞雪无情"
	//b:=[]byte(s)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh))
	fmt.Println(b1)
}

func compulsionTransfer() {
	s := "飞雪无情"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
}

func stringsAndSliceByte() {
	s := "飞雪无情"
	b := []byte(s)
	s3 := string(b)
	fmt.Println(s, string(b), s3)
}

func sliceAndArr() {
	a1 := [2]string{"飞雪无情", "张三"}
	fmt.Printf("函数main数组指针：%p\n", &a1)
	arrayF(a1)
	s1 := a1[0:1]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)
}

// 函数传值后数组打印
func arrayF(a [2]string) {
	fmt.Printf("函数arrayF数组指针：%p\n", &a)
}

// 函数传值后的切片数组打印
func sliceF(s []string) {
	fmt.Printf("函数sliceF Data：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}
