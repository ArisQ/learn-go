package main

import (
	"fibo"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d)=%d\n", i, fibo.Fibonacci(i))
	}
}
