package main

import "fmt"

func main() {
	var result int64 = 0
	for i := 0; i <= 100; i++ {
		result = fibonacci_2(i)
		fmt.Printf("fibonacci_2(%d) is :%d\n", i, result)
	}
}

func fibonacci_2(p int) (res int64) {
	var result *int64 = &res

	if p <= 1 {
		*result = 1
	} else {
		*result = fibonacci_2(p-1) + fibonacci_2(p-2)
	}
	return res
}
