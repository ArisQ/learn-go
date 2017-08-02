package main

import (
	"fmt"
)

func main() {
	fmt.Printf("1 * 2 = %d\n", 1*2)

	aInt := 1
	intPtr := &aInt
	fmt.Printf("intPtr = %p, *intPtr = %d\n", intPtr, *intPtr)
}
