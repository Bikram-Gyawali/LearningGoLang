package main

import "fmt"

func main() {
	defer fmt.Println("\nWorld\n")
	fmt.Println("bikram here")

	myDefer()
}


func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Print(i,"\n")
	}
}