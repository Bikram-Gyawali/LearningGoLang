package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

// go routines practice and mutex in golang

var signals = []string{"test"}
var wg sync.WaitGroup //pointer 

var mut sync.Mutex //pointer

func main() {

	// go greeter("hello")
	// greeter("bikram")

	webSiteList := []string{
		"https://google.com", "https://twitter.com", "https://github.com",
	}

	for _, web := range webSiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("how we doin man")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d statuss code for %s\n", res.StatusCode, endpoint)
	}
}
