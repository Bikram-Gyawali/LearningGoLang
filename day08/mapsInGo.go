package main 

import (
	"fmt"
)

func main(){
	menu:= map[string]float64{
		"momo":150,
		"pizza":200,
		"masu":420,
	}

	fmt.Println(menu)

	fmt.Println(menu["momo"])

	for k,v:=range menu{
		fmt.Println(k,"-",v)
	}

	phoneBook:=map[int]string{
		12342343:"haild",
		42379776:"khali",
		808781897:"asukli",
	}


	fmt.Println(phoneBook)

	fmt.Println(phoneBook[42379776])

	for k,v:=range phoneBook{
		fmt.Println(k,"-",v)
	}

	phoneBook[42379776]="noman"
	fmt.Println(phoneBook)
}

//output 

// map[masu:420 momo:150 pizza:200]
// 150