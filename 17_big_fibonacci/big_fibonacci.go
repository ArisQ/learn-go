package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	i1 := big.NewInt(math.MaxInt64)
	i2 := big.NewInt(65535)
	ir := big.NewInt(1)
	ir.Mul(ir, i1).Add(ir, i2)
	fmt.Println(ir)

	fmt.Println("Fibonacci:")
	f1, f2 := big.NewInt(1), big.NewInt(1)
	fmt.Println(0, f1)
	fmt.Println(1, f2)
	ts := time.Now()
	for i := 2; i < 1000; i++ {
		f1.Add(f1, f2)
		f1, f2 = f2, f1
		fmt.Println(i, f2)
	}
	te := time.Now()
	fmt.Println("Time used:", te.Sub(ts))
}
