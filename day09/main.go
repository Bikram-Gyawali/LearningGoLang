package main

import "fmt"

func main(){
	myNewBill:=newBill("bikramg ko bill")


	
	// fmt.Println(myNewBill)/
	// fmt.Println(myNewBill.format())


	myNewBill.addItem("momo",150)
	myNewBill.updateTip(40)
	fmt.Println(myNewBill.format())

}