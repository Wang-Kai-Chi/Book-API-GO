package handler

import (
	"encoding/json"
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

func MustGetBookFromJson(content string) Book {
	var book Book
	err := json.Unmarshal([]byte(content), &book)
	if err != nil {
		panic(err)
	}
	return book
}

func MustGetBookArrayFromJson(content string) []Book {
	var books []Book
	err := json.Unmarshal([]byte(content), &books)
	if err != nil {
		panic(err)
	}
	return books
}
