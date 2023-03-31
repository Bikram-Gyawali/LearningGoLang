package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hello := "Welcome to our hotel ."

	fmt.Println(hello)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the rating for our pizza :  ")

	//comma or || err err  syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating 10, ", input)

}
