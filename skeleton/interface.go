package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{4.5}
	r1 := Rectangle{5, 7}

	shapes := []Shape{c1, r1}
	for _, shape := range shapes {
		// fmt.Println(shape.area())
		fmt.Println(getArea(shape))
	}
}
