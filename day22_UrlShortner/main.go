package main

import (
	"fmt"
	"net/http"

	"github.com/Bikram-Gyawali/urlshortener/shortenurl"
)


func main()  {
	db,err = shortenurl.Connect()
	if err != nil{
		panic(("failed to coennect database"))
	}
	db.AutoMigrate(&shortenurl.URL{})

	http.HandleFunc("/shorten",func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shorted := shortenurl.ShortenURL(original)

		fmt.Printf(shorted)
		db.Create(&shortenurl.URL{Original: original,Shortened:shorted}),
	})
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.RedirectURL(db,w,r)
	})
	http.ListenANdServe(":5000",nil)
}
