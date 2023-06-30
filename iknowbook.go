package main

import (
	"fmt"

	"iknowbook.com/handler"
)

func main() {
	content := handler.ReadFileAsString("./json/book_single.json")
	fmt.Println(handler.GetBookFromJson(content).PublicationDate)

	ServerStart()
}
