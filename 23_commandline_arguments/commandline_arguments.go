package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var newLine = flag.Bool("n", false, "print with newline")
var times = flag.Int("t", 1, "print the arg t times")

func main() {
	fmt.Println("All arguments:")
	fmt.Println(strings.Join(os.Args, " "))

	fmt.Println("\nflag.PrintDefaults:")
	flag.PrintDefaults()

	fmt.Println("\nParse arguments:")
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			if *newLine {
				s += "\n"
			}
		}
		for j := 0; j < *times; j++ {
			s += flag.Arg(i)
			s += " "
		}
	}
	fmt.Println(s)

	fmt.Println("\nVisit all flags:")
	fmt.Println("|Name|Usage|Value|DefaultValue|")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("|%v|%v|%v|%v|\n", f.Name, f.Usage, f.Value, f.DefValue)
	})
}
