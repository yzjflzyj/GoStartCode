package demo17

func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//双指针
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

//缓存已经计算的结果
var cache = map[int]int{}

func fibForCache(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	result := 0
	switch {
	case n < 0:
		result = 0
	case n == 0:
		result = 0
	case n == 1:
		result = 1
	default:
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}
	cache[n] = result
	return result
}
