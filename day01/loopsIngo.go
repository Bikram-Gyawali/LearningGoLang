package main 

import "fmt"

func main(){
	i:=1
	for i<=5{
		fmt.Printf("b%dkram\n",i)
		i++
	}
	names:=[]string{"hehe","hello","hi","how","hell"}
	
	for index,value :=range names{
		fmt.Printf("the value at index %v is %v\n",index,value)
	}
}

// output ; 
// b1kram
// b2kram
// b3kram
// b4kram
// b5kram

