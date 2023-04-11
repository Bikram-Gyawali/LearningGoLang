package main

import (
	"fmt"
	"sort"
)

func main(){

	fmt.Println("welcome to slices in go lang")

	var fruitList=[]string{"apple","banana","papaya"}

	fmt.Printf("Type of fruitlist %T",fruitList)

	fruitList = append(fruitList, "mango")

	fruitList=append(fruitList[1:])

	highScores:=make([]int,4)

	highScores[0]=234
	highScores[1]=834
	highScores[2]=034
	highScores[3]=264

	fmt.Println(highScores)

	highScores = append(highScores, 055,123,890)

	sort.Ints(highScores)

	var courses=[]string{"reactjs","vue","angular","golang","python","java"}

	var index int = 2

	courses=append(courses[:index],courses[index+1:]...) // remove item form array / slices

}