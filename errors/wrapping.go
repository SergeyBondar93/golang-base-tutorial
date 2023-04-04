package main

import (
	"errors"
	"fmt"
)

var noRows = errors.New("Now rows specified")

func getRecords() error {
	return noRows
}

func webService() error {
	if err := getRecords(); err != nil {
		return fmt.Errorf("Error  --- %w --- when calling DB service", err)
	}
	return nil
}

func main() {
	if err := webService(); err != nil {
		if errors.Is(err, noRows) {
			fmt.Println("The searched record cannot be found", err)
			return
		}
		fmt.Println("unknown error happened during calling db service")
		return
	}

	fmt.Println("Request happened successfully")

}
