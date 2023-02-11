package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}

// PersonSlice 切片，实现sort.Interface
type PersonSlice []Person

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Age < a[i].Age
}

// PersonWrapper
// 动态指定排序规则
type PersonWrapper struct {
	people []Person
	by     func(p, q *Person) bool
}

func (pw PersonWrapper) Len() int { // 重写 Len() 方法
	return len(pw.people)
}
func (pw PersonWrapper) Swap(i, j int) { // 重写 Swap() 方法
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw PersonWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.people[i], &pw.people[j])
}
func main() {
	// 基本升序排序
	baseSort()
	// 基本降序排序
	reverseSort()
	// 通过组合sort.Interface，可以不用重写Len(),Swap(),Less(),也可以根据需要写
	reverseSortWrapper()
	// 根据结构体的属性排序
	sortByStructFiled()
	// 根据结构体的属性排序，动态指定排序规则
	sortByStructFiledAndRules()
	// 优先队列的实现：除了排序接口的实现，还有pop和push方法需要重写，如下hp结构体
}

func sortByStructFiledAndRules() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return q.Age < p.Age // Age 递减排序
	}})

	fmt.Println(people)
	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return p.Name < q.Name // Name 递增排序
	}})

	fmt.Println(people)
}

// sortByStructFiled
//
//	@Description: 按结构体字段排序
func sortByStructFiled() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonSlice(people)) // 按照 Age 的逆序排序
	fmt.Println(people)

	sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的升序排序
	fmt.Println(people)
}

func baseSort() ([]int, []float64, []string) {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)

	//sort.Float64s(doubles) // float64 正序排序 方法1
	//sort.Sort(sort.Float64Slice(doubles))    // float64 正序排序 方法2
	(sort.Float64Slice(float8List)).Sort() // float64 排序方法 方法3

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
	return intList, float8List, stringList
}

func reverseSort() {
	//在 go 中的 sort 包有一个 sort.Interface 接口，该接口有三个方法 Len()、Less()、Swap(i, j)。
	//通用排序函数 sort.Sort 可以排序任何实现了 sort.Interface 接口的对象（变量）。
	//对于 []int、[]float64、[]string 除了使用特殊指定的函数外，
	//还可以使用改装过的类型 IntSlice、Float64Slice 和 StringSlice，然后直接调用它们对应的 Sort() 方法，
	//因为这三种类型也实现了 sortInterface接口，所以可以通过 sort.Reverse 来转换这三种类型的 Interface.Less 方法来实现逆向排序。
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

func reverseSortWrapper() {
	doubles := []float64{3.5, 4.2, 8.9, 100.98, 20.14, 79.32}

	fmt.Printf("doubles is asc ? %v\n", sort.Float64sAreSorted(doubles))
	fmt.Println("after sort by Sort:\t", doubles)

	sort.Sort(Reverse{sort.Float64Slice(doubles)}) // float64 逆序排序
	fmt.Println("after sort by Reversed Sort:\t", doubles)
}

// Reverse
// 自定义的 Reverse 类型
type Reverse struct {
	sort.Interface // 这样， Reverse 可以接纳任何实现了 sort.Interface (包括 Len, Less, Swap 三个方法) 的对象
}

// Less 重写 只是将其中的 Inferface.Less 的顺序对调了一下
func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// 以int切片为例，实现优先队列
type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
