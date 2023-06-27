package main

import (
	"encoding/json"
	"log"
)

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
