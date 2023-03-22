package main

import(
	"fmt"
	"math"
)

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

	var num float64 =12
	var res = math.Sqrt(num)
	var round=math.Round(res)
	fmt.Println(res)
	fmt.Printf("output is %.2f \n",res)
	fmt.Printf("round res is %.2f \n",round)
}