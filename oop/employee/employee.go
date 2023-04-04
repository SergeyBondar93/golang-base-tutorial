package employee

import "fmt"

type employee struct {
	FirstName, LastName string
	Salary              int
}

func (e employee) SayHi() {
	fmt.Println("Hello, my name is ", e.FirstName, " ", e.LastName, " And my salary is ", e.Salary)
}

func New(firstName string, lastName string, salary int) employee {
	e := employee{FirstName: firstName, LastName: lastName, Salary: salary}
	return e
}
