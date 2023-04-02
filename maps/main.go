package main

import "fmt"

type Worker struct {
	firstName, lastName string
}

func updateOriginal(w *Worker) {
	w.firstName = "Changed FIRST NAME"
}
func update(w Worker) {
	w.firstName = "Changed FIRST NAME"
}

func main() {
	emails := make(map[string]string)

	emails["Sergey"] = "sergey@gmail.com"
	emails["Anna"] = "anna@gmail.com"

	// fmt.Println(len(emails))
	// fmt.Println(emails)
	// fmt.Println(emails["Anna"])
	// fmt.Printf("%T\n", emails["Anna"])
	// fmt.Printf("%T\n", emails)

	// delete(emails, "Sergey")

	// fmt.Println(emails)

	// names := map[string]string{"Sergey": "Sergey Bon", "Anna": "Anna Bon"}

	// value, ok := names["Sergey1"]

	// fmt.Println(value, ok)

	Anna := Worker{
		firstName: "Anna",
		lastName:  "Bondar",
	}
	Sergey := Worker{
		firstName: "Sergey",
		lastName:  "Bondar",
	}

	a := 15

	Anna1 := &Anna
	aa := &a

	fmt.Println(aa)

	fmt.Println(Anna1)

	fmt.Println(Anna == Sergey)

	workers := map[Worker]string{
		Anna:   "123456",
		Sergey: "456789",
	}
	fmt.Println(workers)

	workers2 := &workers

	w := *workers2
	w[Anna] = "999999"

	fmt.Println("____UPDATES____")

	fmt.Println(workers)
	fmt.Println(workers2)
	fmt.Println(w)

	// Update the original anna struture
	updateOriginal(&Anna)

	// Update the copy, Anna object will not updated
	update(Anna)

	fmt.Println(w)
	fmt.Println(Anna)

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2, "c": 456}

	fmt.Println(map1 == nil, map2 == nil)

	fmt.Println(compare(map1, map2))

}

func compare(a map[string]int, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	} else {
		for key, value := range a {
			if b[key] != value {
				return false
			}

		}

	}
	return true
}
