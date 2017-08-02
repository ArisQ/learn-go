package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example string"
	runeStr := "This 这 is 是 an 一个 example 实例 string 字符串"
	fmt.Printf("Does \"%s\" have prefix \"%s\"? %t\n", str, "This", strings.HasPrefix(str, "This"))
	fmt.Printf("Does \"%s\" have prefix \"%s\"? %t\n", str, "this", strings.HasPrefix(str, "this"))

	fmt.Printf("Does \"%s\" contains \"%s\"? %t\n", str, "example", strings.Contains(str, "example"))
	fmt.Printf("Does \"%s\" contains \"%s\"? %t\n", str, "Example", strings.Contains(str, "Example"))

	fmt.Printf("The index of \"%s\" in \"%s\" is %d\n", "is", str, strings.Index(str, "is"))
	fmt.Printf("The last index of \"%s\" in \"%s\" is %d\n", "is", str, strings.LastIndex(str, "is"))
	fmt.Printf("The index of \"%s\" in \"%s\" is %d\n", "一", runeStr, strings.Index(runeStr, "一"))
	fmt.Printf("The rune index of \"%s\" in \"%s\" is %d\n", "一", runeStr, strings.IndexRune(runeStr, '一'))

	// Replace
	// Count
	// Repeat
	// ToLower ToUpper
	// TrimeSpace
	// Fields Split
	// Join
	// NewReader Read ReadByte ReadRune

	// strconv package
}
