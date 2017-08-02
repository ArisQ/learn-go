package main

import "fmt"

func main() {
	for i := 5; i >= 0; i-- {
		for j := 5; j >= 0; j-- {
			if r, e := divide(i, j); e != nil {
				fmt.Printf("%d divided by %d error: %e\n", i, j, e)
			} else {
				fmt.Printf("%d / %d = %d\n", i, j, r)
			}
		}
	}

	fmt.Println("\n\nNow unrecoverd panic!")
	i := 0
	_ = i / i
}

func divide(a, b int) (r int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = err.(error)
		}
	}()
	return a / b, nil
}
