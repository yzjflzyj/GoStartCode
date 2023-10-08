package main

import "fmt"

func main() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Println("Before sorting:", arr)
	insertSort(arr)
	fmt.Println("After sorting:", arr)
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		index := i - 1
		insertValue := arr[i]
		for index >= 0 && insertValue < arr[index] {
			arr[index+1] = arr[index]
			index--
		}
		arr[index+1] = insertValue
	}
}
