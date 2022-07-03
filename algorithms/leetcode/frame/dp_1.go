package main

import "fmt"

func main() {
	res := fib(4)
	fmt.Println(res)
}

/**
 *509. 斐波那契数
 */
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a := 1
	c := 0
	for i := 2; i <= n; i++ {
		dpi := a + c
		c = a
		a = dpi
	}
	return a
}
