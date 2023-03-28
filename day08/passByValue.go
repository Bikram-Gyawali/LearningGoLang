package main

import (
	"fmt"
)

func updateName(x string)string{
	x = "ram value changed without pointer"
	return x
}

func updateNameWithPointer(x *string){
	*x = "ram changed with pointer"
}


func updateMenu(y map[string]float64){
	y["tea"]=35
}

func main(){
	name:="bikram"
	
	name =	updateName(name)
	fmt.Println(name)
	
	name1:="bikram1"
	updateNameWithPointer(&name1)
	fmt.Println(name1)


	//maps ma pointer hunxa by default it points to memory locations stored value but mathi hudaina tetti buj simple ma
	menu:= map[string]float64{
		"tea":30,
		"coffee":80,
	}

	updateMenu(menu)
	fmt.Println(menu)

	//value change hunxa  

}