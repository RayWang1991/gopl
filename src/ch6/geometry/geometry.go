package main

import (
	"math"
	"fmt"
)

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	res := float64(0.0)
	for i := range []Point(path) {
		if i > 0 {
			res += path[i-1].Distance(path[i])
		}
	}
	return res
}

func (p *Point) ScaledBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	testPtrMethods()
}

func testPtrMethods() {
	//1
	r := &Point{1, 2}
	r.ScaledBy(2)
	fmt.Println(*r)

	// equivalent to 1
	p := Point{1, 2}
	p.ScaledBy(2)
	fmt.Println(p)

	/*
	// can not take the address (&) to a literal
	// actually the compiler can...
	Point{1,2}.ScaledBy(2)
	*/

	// call a struct method from a pointer is OK, the compiler just obtain the value from the address
	fmt.Println(r.Distance(p))
	fmt.Println((*r).Distance(p))
	// these two are equivalent
	fmt.Println((&Point{1, 3}).Distance(p))
	m := map[*Point]int{}
	m [r] = 1
	fmt.Println(m)
}
