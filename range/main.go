package main

import "fmt"

func main() {
	ids := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	for i := 0; i < len(ids); i++ {
		fmt.Println(ids[i])
	}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for i := 0; i < len(ids); i++ {
		sum += ids[i]
	}
	fmt.Println("SUM: ", sum)

	names := map[string]string{"Sergey": "Sergey Bon", "Anna": "Anna Bon"}

	for k, v := range names {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range names {
		fmt.Println("Name: ", k, " Full Name: ", names[k])
	}

}
