package main

import (
	"flag"
	"os"
	"strconv"
	"fmt"
	"io/ioutil"
	"strings"
)

type Inch float64
type Meter float64

func main() {
	args := flag.Args()
	if len(args) > 0 {
		// collect strings from args
		convStrs(args)
	} else {
		// from the stdIn
		f, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read from std in failed due to %v/n", err)
		} else {
			args := strings.Split(string(f), "\n")
			convStrs(args)
		}
	}
}

func convStrs(strs []string) {
	for _, str := range strs {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert string to float64 failed due to %v/n", err)
		} else {
			conv(num)
		}
	}
}

func conv(num float64) {

}
