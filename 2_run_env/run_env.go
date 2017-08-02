package main

import "fmt"
import "os"
import "runtime"

func main() {
	osName := runtime.GOOS
	fmt.Printf("The operating system is :%s\n", osName)

	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
