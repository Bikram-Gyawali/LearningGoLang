package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadStats(filename string) {
	file, err := os.Open((filename))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Filee Name: %s\n", stats.Name())

	fmt.Printf("Time modified:  %v\n", stats.ModTime().Format("15:04:03"))

}

func ReadWholeFile(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(contents))
}

func ReadByLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()


	scanner :=bufio.NewScanner(file)

	for scanner.Scan(){
		
	}
}

func main() {
	filename := "file.dat"

	// ReadStats(filename)
	ReadWholeFile(filename)
}
