package main

import (
	"fmt"
	"math"
)

// an empty interface the same is any type =  {}interface == any

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

type MyString string

type WithCountA interface {
	CountA() int
}

func counter(obj WithCountA) int {
	return obj.CountA()
}

func (s MyString) CountA() int {
	res := 0

	for _, rune := range s {
		if rune == 'a' {
			res++
		}
	}
	return res
}

func main() {
	circle := Circle{x: 0, y: 0, raduis: 5}
	rectangle := Rectangle{width: 10, height: 10}

	fmt.Println(circle.area())
	fmt.Println(rectangle.area())

	fmt.Println("----------")

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))

	str := MyString("bbbcccaaabbbcccsssaaa") //6

	var byInterface WithCountA

	byInterface = str

	fmt.Println(byInterface.CountA())

	counter(str)

	desc()

}

type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func switcher(n interface{}) string {
	// check type of n
	// like typeof n == someType
	switch v := n.(type) {
	// or just n.(type)
	case string:
		return "This is string"
	case int:
		return "This is int"
	case Circle:
		return "This is Circle - custom type and my Area is " + fmt.Sprintf("%f", v.area())
	case bool:
		return "This is bool"
	default:
		return "I don't know what is it"
	}

}

func desc() {
	p := Person{
		name: "Naveen",
	}
	var w Worker = p
	describe(w)
	w.Work()

	var s interface{} = 56
	assert(s)
	assert("Lol")

	var i int8 = 100
	fmt.Println(switcher(i))
	fmt.Println(switcher(Circle{raduis: 10}))
	fmt.Println(switcher(true))
	fmt.Println(switcher("123"))
	fmt.Println(switcher(123))

}
