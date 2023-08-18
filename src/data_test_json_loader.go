package main

import (
	"encoding/json"

	"iknowbook.com/handler"
)

type DataTestLoader struct {
	getSingleBook    Book
	getBooks         []Book
	getSingleRawBook RawBook
	getRawBooks      []RawBook

	getSingleDvd    Dvd
	getDvds         []Dvd
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

	books := func(raws []RawBook) []Book {
		books := make([]Book, len(raws))
		for i, v := range raws {
			books[i] = singleBook(v)
		}
		return books
	}

	dvd := func(raw RawDvd) Dvd {
		return Dvd{
			Product_: Product{
				Barcode:       raw.Barcode,
				Price:         raw.Price,
				Product_title: raw.Title,
			},
			PublicationDate: raw.PublicationDate,
			Publisher:       raw.Publisher,
		}
	}

	singleRawDvd := func() RawDvd {
		path := "../json/dvd_single.json"
		return mustGetDataFromJson[RawDvd](handler.ReadFileAsString(path))
	}

	dvds := func(raws []RawDvd) []Dvd {
		dvds := make([]Dvd, len(raws))

		for i, v := range raws {
			dvds[i] = dvd(v)
		}
		return dvds
	}

	rawDvds := func() []RawDvd {
		path := "../json/dvd_array.json"
		return mustGetDataFromJson[[]RawDvd](handler.ReadFileAsString(path))
	}

	return DataTestLoader{
		getSingleBook:    singleBook(singleRawBook()),
		getBooks:         books(rawBooks()),
		getSingleRawBook: singleRawBook(),
		getRawBooks:      rawBooks(),

		getSingleDvd:    dvd(singleRawDvd()),
		getDvds:         dvds(rawDvds()),
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
