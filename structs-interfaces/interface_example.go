package main

import (
	"fmt"
	"math"
)

type shapes interface {
	name() string
	area() float64
	perim() float64
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) name() string {
	return "rectangle"
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return 2 * (r.height + r.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) name() string {
	return "circle"
}

func measure(s shapes) {
	fmt.Printf("shape: %s\n", s.name())
	fmt.Println("area: ", s.area())
	fmt.Println("perim: ", s.perim())
}

func interfaceExample() {
	r := rectangle{
		height: 3.0,
		width:  4.0,
	}
	measure(r)

	c := circle{
		radius: 5.0,
	}
	measure(c)
}
