package main

import "fmt"

type Point struct {
	x, y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	test()
}

func test() {
	p := Point{1, 2}
	fmt.Println(p)
	c := Circle{Point{1, 2}, 3}
	fmt.Printf("%v\n", c)
	w := Wheel{Circle{Point{1, 2}, 4}, 10}
	fmt.Printf("%#v\n", w)
}
