package main

import (
	"encoding/json"
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
	PrintBookTitle("json/book_single.json")
}

func PrintBookTitle(jsonFile string) {
	content := ReadFileAsString(jsonFile)

	var books Book
	err := json.Unmarshal([]byte(content), &books)

	if err == nil {
		fmt.Println(books.Title)
	} else {
		fmt.Println(err)
	}
}
