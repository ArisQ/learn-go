package fibo

import "math"

func Fibonacci(i int) int {
	if i == 0 || i == 1 {
		return i
	} else if i > 0 {
		return Fibonacci(i-1) + Fibonacci(i-2)
	} else {
		return Fibonacci(i+2) - Fibonacci(i+1)
	}
}

func Fibonaccif(f float64) float64 {
	return (math.Pow((1+math.Sqrt(5)), f) - math.Pow(1-math.Sqrt(5), f)) /
		math.Pow(2, f) / math.Sqrt(5)
}
