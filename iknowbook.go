package main

import (
	"fmt"
)

type Book struct {
	isbn            string
	title           string
	publicationDate string
	price           string
	author          string
	translator      string
	language        string
}

func main() {
	var fileName = "iknowbook.txt"
	fmt.Println(ReadFileAsString(fileName))
}
