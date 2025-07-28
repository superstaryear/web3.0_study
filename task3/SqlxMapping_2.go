package main

import "fmt"

type Book struct {
	ID     uint
	title  string
	author string
	price  float64
}

func main() {
	var book []Book
	err := DB.Select(&book, "select * from books where price>50")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(book)
}
