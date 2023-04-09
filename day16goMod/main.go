package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("heello ")

	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", servHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r)) 

	fmt.Println("how are yuo")



}

func greet() {
	fmt.Println("hey there goodmorning")
}

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang home</h1>"))
}
