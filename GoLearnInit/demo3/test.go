package main

import "fmt"

func main() {
	// 数组的三种定义方式
	strings1 := [4]string{"a", "b", "c"}
	fmt.Print(strings1)
	strings2 := [...]string{"1", "2", "3"}
	fmt.Print(strings2)
	strings3 := [4]string{0: "x", 2: "y"}
	fmt.Println(strings3)

	// 切片的三种定义方式
	slice1 := []string{"a", "b", "c"}
	fmt.Print(slice1)
	slice2 := make([]string, 4, 8)
	fmt.Print(slice2)
	slice3 := strings3[1:3]
	fmt.Println(slice3)

	// map的两种定义方式
	map1 := map[string]int{"关羽": 12}
	fmt.Println(map1)
	map2 := make(map[string]int)
	map2["123"] = 123
	fmt.Print(map2)
}
