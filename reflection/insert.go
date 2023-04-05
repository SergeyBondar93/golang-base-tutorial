package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId, customerId int
	user                string
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func toType() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}

// the reflect package very useful when we need to check the type of argument

func createQuery(o interface{}) string {
	t := reflect.TypeOf(o)
	k := t.Kind()
	v := reflect.ValueOf(o)
	// it will be a packageName.structureName // e.g main.order
	fmt.Println("Type is ", t)
	// it will be the value of o
	fmt.Println("Value is ", v)

	// struct / int etc
	fmt.Println("Kind of ", k)

	fmt.Println("_____")

	if k == reflect.Struct {
		fmt.Println("Number of fields ", v.NumField())

		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			var normalType string = f.Kind().String()

			// another type check switch
			// switch f.Kind() {
			// case reflect.Array:
			// 	normalType = "Array"
			// case reflect.Int:
			// 	normalType = "Int"
			// case reflect.String:
			// 	normalType = "String"
			// }

			fmt.Printf("Field %d, Value: %v, Reflect Type: %T, NormalType: %v \n", i, f, f, normalType)
		}

	}

	fmt.Println("_____")

	// i := fmt.Sprintf("insert into order values (%d, %d)", o.orderId, o.customerId)
	return "i"
}

func main() {

	o := order{
		orderId:    1234,
		customerId: 567,
	}

	// e := employee{
	// 	name:    "Naveen",
	// 	id:      565,
	// 	address: "Science Park Road, Singapore",
	// 	salary:  90000,
	// 	country: "Singapore",
	// }

	fmt.Println(createQuery(o))

}
