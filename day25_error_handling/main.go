package main

import (
	// "errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

type CustomError struct {
	Message string
	Code    int
}

func (c CustomError) Error() string {
	return c.Message + " " + strconv.Itoa(c.Code)
}

func Divide(x, y int) (float64, error) {
	if y == 0 {
		// return float64(0), errors.New("cannot divide by zero")
		return 0.0, CustomError{"cannot divide by zero", -1}
	} else {
		return float64(x) / float64(y), nil
	}

}

func main() {
	file, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("error reading file", err)
	} else {
		fmt.Println(string(file))
	}

	value, divErr := Divide(7, 0)

	if divErr != nil {
		fmt.Println(divErr)
	} else {
		fmt.Println(value)
	}

}
