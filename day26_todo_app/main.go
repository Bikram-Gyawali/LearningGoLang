package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Learn golang", Done: false},
			{Item: "Make gf", Done: false},
			{Item: "Sleep on time", Done: false},
			{Item: "Overthinking", Done: true},
			{Item: "Play chess", Done: true},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":3001", mux))
}
