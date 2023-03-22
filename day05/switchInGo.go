package main

import( 
	"fmt"
	"time"
)

func main(){
	num:=5

	// if num < 5 {
	// 	fmt.Println("Hello ",num)
	// }else if num > 10{
	// 	fmt.Println("bye",num)
	// }else{
	// 	fmt.Println("hehehebhbhe")
	// }

	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Def")
	}


	fmt.Println("When is saturday")
	
	today:= time.Now().Weekday()
	switch time.Saturday{
	case today+0:
		fmt.Println("Todaay")
	case today+1:
		fmt.Println("Tomorrow")
	case today+2:
		fmt.Println("Parsi")
	default:
		fmt.Println("alli badi din lagxa")
	}
}