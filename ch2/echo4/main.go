package main

import (
	"fmt"
	"flag"
	"strings"
)

var n = flag.Bool("n", false, "omit the trailing new line")
var sep = flag.String("s", " ", "seperator string")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
