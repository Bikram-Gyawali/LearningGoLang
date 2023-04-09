package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to golang api development")

	PerformGetRequst()

	PerformPostJsonRequest()
}

func PerformGetRequst() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status codee : ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCOunt, _ := responseString.Write(content)

	fmt.Println("Byte COunt is", byteCOunt)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"coursename":"Lets go learn golang",
			"price":0,
			"platform":"learnCodeOnline.com",
		}
	`)

	response, err := http.Post(myUrl, "application/jsosn", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(content)
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postForm"

	data := url.Values{}

	data.Add("firstName", "bikram")
	data.Add("lastname", "gyawali")
	data.Add("age", "20")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
