package main

import (
	"fmt"
	"reflect"
	"strconv"
	"structs/computer"
	"structs/computer/subfolder"
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

func print(p Person) {
	fmt.Println("__PERSON_START__")
	v := reflect.ValueOf(p)

	num := v.NumField()

	for i := 0; i < num; i++ {
		name := v.Type().Field(i).Name
		fmt.Println(name + " - " + v.FieldByName(name).String())

	}

	fmt.Println("__PERSON_END__")
}

type AnonFields struct {
	string
	int
}

func main() {
	person1 := Person{firstName: "Sergey", lastName: "Bondar", city: "Hamburg", age: 29, gender: "m"}
	person2 := Person{firstName: "Anna", lastName: "Riffel", city: "Hamburg", age: 33, gender: "f"}

	person3 := &Person{"Lol", "Kek", "Hamburg", "f", 77}
	fmt.Println(person3)
	print(person1)
	print(person2)
	// print(person3) // to pass a pointer to a struct to function , we should use a * symbol in func to access to the value of struct

	fmt.Println(person3)

	person1.firstName = "SERGEY"

	fmt.Println(person1)
	// /the next couple of lines is equals
	person1.hasBirthday()
	(&person1).hasBirthday()

	fmt.Println("AFTER UPDATE AGE", person1)

	// fmt.Println(person1.greet())

	fmt.Println(person2)
	person2.getMarried("Bondar")
	fmt.Println(person2)

	anon := AnonFields{"Lol", 132}

	fmt.Println(anon)
	fmt.Println(anon.int)

	nastedStructs()
	promotedFields()

	importedStructs()

}

type Internal struct {
	deepField string
}
type External struct {
	common         string
	internalStruct Internal
}

func nastedStructs() {
	deep := External{
		common: "abc",
		internalStruct: Internal{
			deepField: "Deep field",
		},
	}

	fmt.Println(deep.internalStruct.deepField)

}

type WithPromotedFields struct {
	one, two, three string
	PromotedFields
}
type PromotedFields struct {
	four, five string
}

func promotedFields() {
	promoted := WithPromotedFields{one: "one - 1", two: "two - 2", three: "three - 3", PromotedFields: PromotedFields{four: "four - 4", five: "five - 5"}}

	fmt.Println(promoted)

	fmt.Println(promoted.one)
	fmt.Println(promoted.two)
	fmt.Println(promoted.three)
	fmt.Println(promoted.four)
	fmt.Println(promoted.five)
}

func importedStructs() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
	}

	fromSubfolder := subfolder.SubspecComp{ThisIsSubspec: "This is struct from subfolder of computer package"}

	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
	fmt.Println("Subspec:", fromSubfolder.ThisIsSubspec)
}
