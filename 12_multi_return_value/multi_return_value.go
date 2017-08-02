package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	a, b := 1, 2
	s, p, d := SumProductDiffNamed(a, b)
	fmt.Printf("The named sum,product,difference of %d and %d is %d,%d,%d\n", a, b, s, p, d)

	x := 0.1
	sx, err := MySqrtNamed(x)
	fmt.Printf("The named squre root of %f is %f, error:%e\n", x, sx, err)
	sx, err = MySqrtUnnamed(x)
	fmt.Printf("The unnamed squre root of %f is %f, error:%e\n", x, sx, err)
	x = -0.1
	sx, err = MySqrtNamed(x)
	fmt.Printf("The named squre root of %f is %f, error:%e\n", x, sx, err)
	sx, err = MySqrtUnnamed(x)
	fmt.Printf("The unnamed squre root of %f is %f, error:%e\n", x, sx, err)
}

func SumProductDiffNamed(a, b int) (sum, product, diff int) {
	sum = a + b
	product = a * b
	diff = a - b
	return
}

func SumProductDiffUnnamed(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func MySqrtNamed(x float64) (sx float64, err error) {
	if x < 0 {
		err = errors.New("x小于0")
	} else {
		sx = math.Sqrt(x)
	}
	return
}

func MySqrtUnnamed(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("x小于0")
	} else {
		return math.Sqrt(x), nil
	}
}
