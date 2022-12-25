package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//集合
	//数组
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])
	//省略长度,只适用于所有元素都被初始化的情况
	array1 := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(array1)
	//初始化部分元素
	array2 := [5]string{1: "b", 3: "d"}
	fmt.Println(array2)

	//数组遍历
	for i, v := range array {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}

	//切片  基于数组生成切片，包含索引start，但是不包含索引end:  slice:=array[start:end]
	slice := array[2:5]
	fmt.Println(slice)
	//可省略start或end,或都省略 array[:] 等价于 array[0:5]。
	slice[1] = "f"
	fmt.Println(array)

	//切片声明:make函数:slice1:=make([]string,4,8),容量可省略
	slice1 := make([]string, 4, 8)
	fmt.Println(slice1)
	//字面量声明
	slice2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice2), cap(slice2))

	//append函数
	//追加一个元素
	slice2 = append(slice, "f")
	//多加多个元素
	slice2 = append(slice1, "f", "g")
	//追加另一个切片
	slice2 = append(slice1, slice...)
	fmt.Println(slice2)
	for i, v := range slice2 {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}

	//map
	//1.make函数声明
	nameAgeMap := make(map[string]int)
	nameAgeMap["关羽"] = 20
	fmt.Println(nameAgeMap)
	//2.字面量声明:
	nameAgeMap1 := map[string]int{"张飞": 20}
	fmt.Println(nameAgeMap1)

	//获取和删除
	//添加键值对或者更新对应 Key 的 Value
	nameAgeMap["飞雪无情"] = 18
	//获取指定 Key 对应的 Value
	age := nameAgeMap["飞雪无情"]
	fmt.Println(age)
	//如果 Key 不存在，返回的 Value 是该类型的零值,因此一般需要先判断key的存在
	age, ok := nameAgeMap["飞雪无情1"]
	if ok {
		fmt.Println(age)
	}
	// 删除
	delete(nameAgeMap, "飞雪无情")

	//遍历
	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}

	//String 和 []byte
	s := "Hello飞雪无情" //字节长度17
	bs := []byte(s)
	c := []rune(s) //Unicode数组
	fmt.Println(c, bs, s, len(s))
	//utf8.RuneCountInString展示字符长度
	fmt.Println(utf8.RuneCountInString(s))
	//按Unicode编码遍历
	for i, r := range s {
		fmt.Println(i, r)
		fmt.Printf("Unicode: %c  %d\n", r, r)
	}
}
