package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "This is a string.\n这是一个字符串。"
	fmt.Println("The original string:\n", str)
	str2 := strings.Map(func(r rune) rune {
		if r < utf8.RuneSelf {
			return r
		}
		return '?'
	}, str)
	fmt.Println("The string after map:\n", str2)
}
