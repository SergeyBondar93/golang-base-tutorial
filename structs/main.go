package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

type Person2 struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)

func (person Person) greet() string {
	return "Hello, my name is " +
		person.firstName +
		" " +
		person.lastName +
		" and I am from " +
		person.city +
		" and I am " +
		strconv.Itoa(person.age)
}

/*
current value reciver *
* - get access to original value of created structure. this this in JavaScript
without * - copy of this structure
*/
func (p *Person) hasBirthday() {
	fmt.Println("!!!", p)
	p.age++
}

// getMarried
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Sergey", lastName: "Bondar", city: "Hamburg", age: 29, gender: "m"}
	person2 := Person{firstName: "Anna", lastName: "Riffel", city: "Hamburg", age: 33, gender: "f"}

	person1.firstName = "SERGEY"

	fmt.Println(person1)
	person1.hasBirthday()
	fmt.Println(person1)

	// fmt.Println(person1.greet())

	fmt.Println(person2)
	person2.getMarried("Bondar")
	fmt.Println(person2)

}
