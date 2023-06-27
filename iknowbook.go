package main

import (
	"fmt"
)

type Book struct {
	Isbn            string
	Title           string
	PublicationDate string
	Price           string
	Author          string
	Translator      string
	Language        string
}

func main() {
	fmt.Println(GetBookArrayFromJson("json/book_arr.json")[4].Title)
}
