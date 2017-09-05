package tempconv

import (
	"flag"
	"os"
	"strconv"
	"fmt"
	"io/ioutil"
	"strings"
)

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
	// show C,F,K,I,M
	c := Celsius(num)
	fmt.Printf("C:%s F:%s K:%s\n", c, CToF(c), CToK(c))
	f := Fahrenheit(num)
	fmt.Printf("F:%s C:%s K:%s\n", f, FToC(f), FToK(f))
	k := Kelvin(num)
	fmt.Printf("K:%s C:%s F:%s\n", k, KToC(k), KToF(k))
	i := Inch(num)
	fmt.Printf("In:%s M:%s \n", i, InToMeter(i))
	m := Meter(num)
	fmt.Printf("M:%s In:%s \n", m, MeterToInch(m))
}
