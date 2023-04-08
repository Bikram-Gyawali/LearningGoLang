package main

import (
	"fmt"
	"net/url"
	"strings"
)

const myurl string = "https://localhost:3000/study?q=4214sa&id=dsmi13&name=bikram&course=golang&position=developer&position=fullstack"

func main() {

	fmt.Println("Welcomoe to handling urs in golang")

	fmt.Println(strings.Split(myurl, "/"))

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawPath)
	fmt.Println(result.RawQuery)
	fmt.Println(result.RawFragment)

	qparams := result.Query()
	fmt.Printf("the type of querey params are : %T\n", qparams)

	fmt.Println("The query params are: \n", qparams)

	fmt.Println("The value of position are: \n", qparams["position"])

	//crete url using parts of url
	partsOfUrl:=&url.URL{
		Scheme:"https",
		Host:"lcalhost",
		Path:"/study",
		RawPath: "name=bikram",
	}

	urlConcat:=partsOfUrl.String() 

	fmt.Println(urlConcat)


}
