package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int8
}

func (p Person) Describe() {
	fmt.Printf("Hello I am %s  and I am %d \n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Println("I Am just a City named " + a.state + " in " + a.country + "\n")
}

func main() {
	var d1 Describer
	p1 := Person{"Sergey", 29}
	d1 = p1
	d1.Describe()

	p2 := Person{"Anna", 33}
	d1 = &p2

	d1.Describe()

	var d2 Describer

	a := Address{"Hamburg", "Germany"}

	// this does not working cuz Address has a Describe method on pointer reviever
	// and it does not exist on the structure, only on the instance
	// d2 = a

	d2 = &a

	d2.Describe()
	a.Describe()
	(&a).Describe()
}
