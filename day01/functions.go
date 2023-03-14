package main

import (
    "fmt"
    "math"
)

func main(){
	num1:=10
	num2:=32

	a:=2.0
	b:=5.0

	fmt.Println("sum ", sum(num1,num2))  
	fmt.Println("subctraction",sub(num1,num2))
	fmt.Println("power",power(a,b))
	
	// var result1 int
	// var result2 int

	result1,result2:=calc(num1,num2)

	fmt.Println(result1,result2)
}


func sum(a int,b int)int{
	return a + b
}

func sub(a int , b int) int {
	return a -b 
}

func power(a float64,b float64)float64 {
	return math.Pow(a,b)
}

func calc(a int,b int )(int, int){ // here in the parenthesis there are two int so that the function returns 2 values 
 	var mul=a*b
	var div=b/a
	return mul,div
}


// output 
// sum 42
// sub -22