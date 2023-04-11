package main

import "fmt"

func main() {
	fmt.Println("if else in golang")
	a := 10
	var result string

	if a < 10 {
		result = "Not less than 10"
	} else if a > 10 {
		result = "greater than 10"
	} else {
		result = "equal to 10"
	}

	fmt.Println(result)
}
