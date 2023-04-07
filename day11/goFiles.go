package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go lang")
	content := "This needs to go in a file"

	file, err := os.Create("./myGoFile")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)


	length, err := io.WriteString(file, content)
	checkNilErr(err)


	fmt.Println("lenth is : ", length)

	defer file.Close()
	readFile("./myGoFile")

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("data in file is \n", string(databyte))
}


func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}