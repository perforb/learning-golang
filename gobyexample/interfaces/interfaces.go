package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return r.width*2 + r.height*2
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{
		width:  3,
		height: 4,
	}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
