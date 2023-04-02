package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, raduis float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.raduis * c.raduis
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, raduis: 5}
	rectangle := Rectangle{width: 10, height: 10}

	fmt.Println(circle.area())
	fmt.Println(rectangle.area())

	fmt.Println("----------")

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))

}
