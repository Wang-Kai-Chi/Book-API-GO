package main

import (
	"fmt"

	"iknowbook.com/handler"
)

func main() {
	fmt.Println(handler.GetBookFromJson("./json/book_single.json").PublicationDate)
}
