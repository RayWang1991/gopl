package main

import (
	"fmt"
	"flag"
)

type unit int

const (
	cel unit = iota
	fah
	kel
)

type tempflag struct {
	v float64
	u unit
}

var t = TempFlag(20, cel, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(t)
}

// register a flag named `temp`
func TempFlag(value float64, u unit, usage string) *tempflag {
	t := tempflag{value, u}
	flag.CommandLine.Var(&t, "temp", usage)
	return &t
}

func (t *tempflag) String() string {
	switch t.u {
	case cel:
		return fmt.Sprintf("%.2fC",t.v)
	case fah:
		return fmt.Sprintf("%.2fF",t.v)
	case kel:
		return fmt.Sprintf("%.2fK",t.v)
	}
	return ""
}

func (t *tempflag) Set(s string) error {
	// default is C
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "c":
		t.u = cel
		t.v = value
		return nil
	case "F", "f":
		t.u = fah
		t.v = value
		return nil
	case "K", "k":
		t.u = kel
		t.v = value
		return nil
	}
	return fmt.Errorf("invalid temperatue %q", s)
}
