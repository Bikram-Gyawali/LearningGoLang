package main

import "fmt"


func main(){
	fmt.Println("Maps in go lang")

	languages:= make(map[string]string)

	languages["JS"]="Javascript"
	languages["PY"]="Python"
	languages["c++"]="C++"
	languages["go"]="golang"

	fmt.Println("JS short form",languages["JS"])

	delete(languages,"c++") //remove item from map

	//loops in maps
	for k,v :=range languages{
		fmt.Printf("For key %v value is %v \n",k,v)
	}
}