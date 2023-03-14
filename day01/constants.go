package main

import "fmt"

func main(){
	const val1=40
	const val2=20
	const sum=val1+val2;
	const multiplication=val1*val2
	const div=val2/val1
	fmt.Println(sum,multiplication,div)
	// val1=30 gives error ./constants.go:12:2: cannot assign to val1 (untyped int constant 40)
}