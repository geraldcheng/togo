package main

import "fmt"
import "time"

const LIM = 41

func main() {
	start := time.Now()
	result := 0
	for i := 0; i <= LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is :%d location is:%d\n", i, result, i)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
