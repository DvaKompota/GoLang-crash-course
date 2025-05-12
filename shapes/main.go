package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{base: 3, height: 4}
	s := square{sideLength: 2}
	printArea(t)
	printArea(s)
}