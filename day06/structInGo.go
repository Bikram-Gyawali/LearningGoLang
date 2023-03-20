package main

import "fmt"

type Student struct{
	rollno int
	name string
	marks int
}


func main(){
	var s1= Student{3,"Bikram",79}
	// var s1 Student= Student{3,"Bikram",79} we can also write like this


	fmt.Println(s1)

	fmt.Println(s1.name)

	var s2= Student{rollno:88,name:"ram"}
	fmt.Println(s2)
}