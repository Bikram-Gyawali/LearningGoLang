package main

import (
	"fmt"
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

func main() {
	filename := "file.dat"

	ReadStats(filename)
}
