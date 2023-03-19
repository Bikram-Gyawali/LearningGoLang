package main

import "fmt"


func main(){
a()

defer fmt.Println("hello defer")
for i:=0;i<10;i++ {
	defer fmt.Println(i);
}
}

func a(){
	fmt.Println("a called")
	defer b()
	fmt.Println("a ends")
}

func b(){
	fmt.Println("in b")
}