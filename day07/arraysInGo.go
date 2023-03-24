package main

import "fmt"


func main(){
	var ages [3]int = [3]int{10,20,13}

	// shorthand
	var agesS=[3]int{10,20,12}

	fmt.Println(ages,len(agesS))

	names:=[]string{"he","is","a","good","boy"}
	names[1]="is not"

	fmt.Println(names,len(names))

	var points =[]int{100,80,53,43,89,52}
	points[0]=10
	points=append(points,66)

	fmt.Println(points,len(points))


	// slice ranges
	rangeOne:=names[1:3]

	rangeTwo:=names[1:]

	fmt.Println("range1",rangeOne)

	fmt.Println("range2",rangeTwo)
} 