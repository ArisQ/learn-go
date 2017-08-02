package fibo

import (
	"testing"
)

var testData = []struct {
	index     int
	fiboValue int
}{
	// Normal data
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
}

var testNegativeData = []struct {
	index     int
	fiboValue int
}{
	// Negative data
	{-1, 1},
	{-2, -1},
	{-3, 2},
	{-4, -3},
	{-5, 5},
	{-6, -8},
	{-7, 13},
	{-8, -21},
}

func TestFibonacci(t *testing.T) {
	for _, d := range testData {
		v := Fibonacci(d.index)
		if v != d.fiboValue {
			t.Fatalf("Fibonacci(%d)!=%d,which is %d", d.index, d.fiboValue, v)
		}
	}
	for _, d := range testNegativeData {
		v := Fibonacci(d.index)
		if v != d.fiboValue {
			t.Fatalf("Fibonacci(%d)!=%d,which is %d", d.index, d.fiboValue, v)
		}
	}
}

func TestFibonaccif(t *testing.T) {
	for _, d := range testData {
		v := int(Fibonaccif(float64(d.index)))
		if v != d.fiboValue {
			t.Fatalf("Fibonaccif(%d)!=%d,which is %d", d.index, d.fiboValue, v)
		}
	}
}

func BenchmarkFibonacci10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonaccif10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonaccif(10)
	}
}

func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkFibonaccif20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonaccif(20)
	}
}
