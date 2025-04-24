package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.height + r.width
}

func (c circle) area() float64 {
	return c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perimeter())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle with radius: ", c.radius)
	}
}

func geometrySwitch(g geometry) {
	switch g.(type) {
	case circle:
		fmt.Println("Circle with radius: ", g.(circle).radius)
	case rect:
		fmt.Println("Rectangle with width: ", g.(rect).width, ", height: ", g.(rect).height)
	default:
		fmt.Println("Not a geometry!")
	}
}

func main() {
	var r rect = rect{width: 3, height: 4}
	c := circle{radius: 4}
	measure(r)
	measure(c)
	detectCircle(c)
	geometrySwitch(r)
	geometrySwitch(c)
}
