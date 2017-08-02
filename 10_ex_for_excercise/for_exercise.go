package main

import "fmt"

func main() {
	forLoop()
	forCharacter()
	bitwiseComplement()
	fizzBuzz()
	rectangleStars()
}

func forLoop() {
	//for loop
	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	//goto
	i := 0
LABEL:
	fmt.Printf("%d ", i)
	i++
	if i < 15 {
		goto LABEL
	}
	fmt.Println()
}

func forCharacter() {
	//two layer loop
	for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	//one layer loop
	str := "GGGGGGGGGGGGGGGGGGGGGGGGG"
	for i := 1; i <= len(str); i++ {
		fmt.Printf(str[:i])
		fmt.Println()
	}
}

func bitwiseComplement() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("bitwise of %d is %b\n", i, ^i)
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
}

func rectangleStars() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
