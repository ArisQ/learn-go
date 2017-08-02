package main

import (
	"fmt"
	"runtime"
	"stack"
)

func main() {
	s := stack.Stack{}
	for i := 0; i < 4; i++ {
		fmt.Println("pushing", i)
		s.Push(i)
		fmt.Println(s.String())
	}
	for i := 0; i < 5; i++ {
		s.Pop()
		fmt.Println(s.String())
	}
	fmt.Println(s.String())
	stackFinalizerTest()
	runtime.GC()
}

func stackFinalizerTest() {
	s := stack.NewStack()
	fmt.Println("NewStack")
	for i := 0; i < 4; i++ {
		fmt.Println("pushing", i)
		s.Push(i)
		fmt.Println(s.String())
	}
	for i := 0; i < 5; i++ {
		s.Pop()
		fmt.Println(s.String())
	}
	fmt.Println(s.String())
}
