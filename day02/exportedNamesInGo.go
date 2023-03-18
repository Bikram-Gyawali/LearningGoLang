package main

import "fmt"

func main(){
	fmt.Println("Bikraam")
	// fmt.toLarge(10)
	Bemo()
	demo()
}

// here Bemo function is accessible outside the package as its name's first letter is capital
func Bemo(){
	fmt.Println("To access any value or function it should be named the first letter capital")
}

func demo(){
	fmt.Println("This function is not accessible outside the package")
}