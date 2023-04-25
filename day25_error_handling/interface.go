package main

import "fmt"

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

type printer interface {
	printInfo()
}

func empty(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I ama ring :  %s\n", i)
	case int:
		fmt.Printf("I am an int %d\n", i)

	case Book:
		fmt.Printf("I ama a book %s", i.(Book).Title)
	}
}

func main() {
	gunslinger := Book{
		Title: "hello world",
		Price: 1.23,
	}

	tea := Drink{
		Name:  "Masala tea",
		Price: 30,
	}

	gunslinger.printInfo()
	tea.printInfo()

	info := []printer{gunslinger, tea}

	info[0].printInfo()
	info[1].printInfo()

	empty("bikram")
	empty(234)
	empty(gunslinger)
}

func (d Drink) printInfo() {
	fmt.Printf("Drink : %s Price: %.2f\n", d.Name, d.Price)
}

func (d Book) printInfo() {
	fmt.Printf("Book : %s Price: %.2f\n", d.Title, d.Price)
}
