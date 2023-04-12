package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	fmt.Println("Switch and case dice gamee")

	rand.Seed(time.Now().UnixNano())
	diceNUmber := rand.Intn(6)+1

	fmt.Println("Value of dice is : ", diceNUmber)


	switch diceNUmber{
	case 1:
		fmt.Println("Dice value is one ")
	case 2:
		fmt.Println("Move 2 spot")
	case 3:
		fmt.Println("move 3 spot")
	case 4 :
		fmt.Println("move 4 spot")
	case 5:
		fmt.Println("move 5 spot")
	case 6:
		fmt.Println("mvoee 6 spot and roll dice again")
	default:
		fmt.Println("this is not availablef")
	}
}