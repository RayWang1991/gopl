package main

import (
	"gopl/src/ch2/tempconv"
	"fmt"
	"flag"
)

type unit int

const (
	Cel = iota
	Fah
	Kel
)

type tempflag struct {
	tempconv.Celsius
	//tempconv.Fahrenheit
	//tempconv.Kelvin
}

var t = TempFlag1(20.00,"temperatue")

// register a flag named `temp`
func TempFlag1(value float64, usage string) *tempflag {
	temp := tempflag{tempconv.Celsius(value)}
	flag.CommandLine.Var(&temp, "temp", usage)
	return &temp
}

func main() {
	fmt.Println(t)
	flag.Parse()
}

func (t *tempflag) Set(s string) error{
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "c":
		t.Celsius = tempconv.Celsius(value)
	//	t.Fahrenheit = tempconv.CToF(t.Celsius)
	//	t.Kelvin = tempconv.CToK(t.Celsius)
	//	return nil
	//case "F", "f":
	//	t.Fahrenheit = tempconv.Fahrenheit(value)
	//	t.Celsius = tempconv.FToC(t.Fahrenheit)
	//	t.Kelvin = tempconv.FToK(t.Fahrenheit)
	//	return nil
	//case "K", "k":
	//	t.Kelvin = tempconv.Kelvin(value)
	//	t.Celsius = tempconv.KToC(t.Kelvin)
	//	t.Fahrenheit = tempconv.KToF(t.Kelvin)
		return nil
	}
	return fmt.Errorf("Invalid temperatue")
}
