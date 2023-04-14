package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByWord(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByBytes(filename string, size uint8) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, size)

	for {
		totalRead, err := file.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}

		fmt.Println(buf[:totalRead])
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "=")
		fmt.Println(raw[0])
		fmt.Println(raw[1])
	}
}

func main() {
	// filename := "file.dat"
	filenameCfg:= "test.cfg"
	// ReadStats(filename)
	// ReadWholeFile(filename)
	// ReadByWord(filename)
	// ReadByLine(filename)
	ReadByBytes(filenameCfg)
}
