package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type %T\n", response)

	defer response.Body.Close() // my responsibility to close th connection

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)

}
