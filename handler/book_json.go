package handler

import (
	"encoding/json"
	"log"
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

func GetBookFromJson(jsonFile string) Book {
	content := ReadFile(jsonFile)

	var book Book
	err := json.Unmarshal(content, &book)

	if err != nil {
		log.Fatal(err)
	}
	return book
}

func GetBookArrayFromJson(jsonFile string) []Book {
	content := ReadFile(jsonFile)

	var books []Book
	err := json.Unmarshal(content, &books)

	if err != nil {
		log.Fatal(err)
	}
	return books
}
