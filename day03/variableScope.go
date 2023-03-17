package main

import "fmt"

var b=15
var a=11   // we call it package level variable

func calc(){
	var a = 10 // we call it function level variable
	fmt.Println("in calc local",a) 
}

func main(){
	calc()
	// fmt.Println(a) -->  returns undefined a error
	fmt.Println("in main a",a);
	fmt.Println(b)
}