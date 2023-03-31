package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main(){
	fmt.Println("welcome to out pizza app")
	fmt.Println("\n please rate our pizza 1-5")

	reader:=bufio.NewReader(os.Stdin)


	input , _ := reader.ReadString('\n') // read until \n or enter key 

	fmt.Println("Thanks for rating,", input)


	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("ADDED 1 TO YOUR RATING: ", numRating+1)
	}
}