package main

import "sort"

func main() {
	data := []int{2, 4, 1, 5, 3, 8, 35, 23, 78}
	x := 100
	i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
	if i < len(data) && data[i] == x {
		// x is present at data[i]
	} else {
		// x is not present in data,
		// but i is the index where it would be inserted.
	}

}
