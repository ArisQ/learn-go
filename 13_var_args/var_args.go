package main

import "fmt"

func main() {
	printVarArgs(1, 2, 3)

	arg := []int{5, 6, 7}
	fmt.Println("print of ", arg)
	printVarArgs(arg...)
}

func printVarArgs(arg ...int) {
	for _, v := range arg {
		fmt.Println(v)
	}
}
