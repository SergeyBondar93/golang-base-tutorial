package main

import (
	"errors"
	"fmt"
	"math"
)

type areaError struct {
	Err    string
	Radius float64
}

func (e *areaError) IsNegative() bool {
	return e.Radius < 0
}

func (e *areaError) Error() string {
	return e.Err
	// switch e.Err {
	// case "Radius is negative":
	// 	return fmt.Sprintf("You cannot provide radius less then zero. Provided %0.2f", e.Radius)
	// case "Radius is zero":
	// 	return fmt.Sprintf("You cannot provide radius equals to zero. Provided %0.2f", e.Radius)
	// default:
	// 	return fmt.Sprintf("Unhandled error")
	// }
}

func New(text string, radius float64) error {
	return &areaError{text, radius}
}

func circleArea(radius float64) (float64, error) {
	if radius == 0 {
		return 0, &areaError{Err: "Radius is zero", Radius: radius}
	}
	if radius < 0 {
		return 0, &areaError{Err: "Radius is negative", Radius: radius}
	}
	return radius * radius * math.Pi, nil
}

// direct comparison, the easiest way to compare errors
// errors.Is(ourError, someErrorFromPackage)

func main() {
	radius := -0.40

	area, err := circleArea(radius)

	if err != nil {
		// is next comparison is true,
		// we should use this error object to get more info about error
		var areaError *areaError
		if errors.As(err, &areaError) {
			fmt.Println(areaError.Radius, "My custom error happened <3", err)
		} else {
			fmt.Println("General error happened", err)
		}

	} else {
		fmt.Println("Radius is ", area)
	}

}
