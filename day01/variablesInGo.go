package main

import "fmt"

func main(){

	//var values // we cannot declare a variable without any types but we can assign a variable withoud give its types as below
	// var values1=15 

	// var val3 int
	// val3 = 30

	// go compiler will give errors as below if the variables are declared and not used
// 	./variablesInGo.go:8:6: values1 declared and not used
// ./variablesInGo.go:10:6: val3 declared and not used

	//declaring and assigning value 
	var val1 int =10 // go has different types the int supports entire range and the uint supports only positive values and we cannot perform calculation with variables with diffrent types
	var val2 int =20
	var sum=val1+val2;
	var multiplication=val1*val2
	var div=val2/val1
	fmt.Println(sum,multiplication,div)


	//declaring 2 variables in same line 
	var smline1,smline2 int

	//assigning 2 values in same line
	smline1,smline2 = 44,55

	fmt.Println(smline1,smline2) // 44 , 55


	// the variables can also be declared without writing the var keyword in shorthand 
	shorthandVar:=10
	fmt.Println(shorthandVar)
}