package main

import (
	"encoding/json"

	"iknowbook.com/handler"
)

type DataTestLoader struct {
	getSingleBook    Book
	getSingleRawBook RawBook
	getRawBooks      []RawBook

	getSingleRawDvd RawDvd
	getRawDvds      []RawDvd
}

func NewDataTestLoader() DataTestLoader {
	singleRawBook := func() RawBook {
		path := "../json/book_single.json"
		return mustGetDataFromJson[RawBook](handler.ReadFileAsString(path))
	}

	singleBook := func(raw RawBook) Book {
		pd := Product{
			Barcode:       raw.Isbn,
			Price:         raw.Price,
			Product_title: raw.Title,
		}

		book := Book{
			Product_:        pd,
			PublicationDate: raw.PublicationDate,
			Author:          raw.Author,
			Translator:      raw.Translator,
			Language:        raw.Language,
		}
		return book
	}

	rawBooks := func() []RawBook {
		path := "../json/book_arr.json"
		return mustGetDataFromJson[[]RawBook](handler.ReadFileAsString(path))
	}

	singleRawDvd := func() RawDvd {
		path := "../json/dvd_single.json"
		return mustGetDataFromJson[RawDvd](handler.ReadFileAsString(path))
	}

	rawDvds := func() []RawDvd {
		path := "../json/dvd_array.json"
		return mustGetDataFromJson[[]RawDvd](handler.ReadFileAsString(path))
	}

	return DataTestLoader{
		getSingleBook:    singleBook(singleRawBook()),
		getSingleRawBook: singleRawBook(),
		getRawBooks:      rawBooks(),

		getSingleRawDvd: singleRawDvd(),
		getRawDvds:      rawDvds(),
	}
}

func mustGetDataFromJson[T Data](content string) T {
	var entity T
	err := json.Unmarshal([]byte(content), &entity)
	if err != nil {
		panic(err)
	}
	return entity
}
