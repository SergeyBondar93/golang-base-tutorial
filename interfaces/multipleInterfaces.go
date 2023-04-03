package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type CommonInterface interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	name                                   string
	basicPay, pf, totalLeaves, leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("I am %s and my salary is %d", e.name, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{name: "Sergey Bondar", basicPay: 100, pf: 50, totalLeaves: 30, leavesTaken: 5}

	//only one method DisplaySalary
	var s SalaryCalculator = e

	s.DisplaySalary()

	//only one method CalculateLeavesLeft
	var l LeaveCalculator = e
	fmt.Println("\n Leaves left = ", l.CalculateLeavesLeft())

	// both methods are available
	var c CommonInterface = e
	c.CalculateLeavesLeft()
	c.DisplaySalary()

	// all the methods and fields available
	var aaa = e
	aaa.CalculateLeavesLeft()
	aaa.DisplaySalary()
}
