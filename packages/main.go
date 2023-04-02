package main

import (
	"fmt"
	"math"

	"github.com/sergeyBondar93/golang-base-tutorial/packages/strutil"
	"github.com/sergeyBondar93/golang-base-tutorial/packages/times"
)

func main() {
	fmt.Println(math.Floor(2.8))
	fmt.Println(math.Abs(-2.8))
	fmt.Println(math.Ceil(2.8))
	fmt.Println(strutil.Reverse("This string will be reversed"))
	fmt.Println(times.Times(2, 3))
}
