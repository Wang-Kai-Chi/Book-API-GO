package main

import (
	"encoding/json"

	"iknowbook.com/handler"
)

type RawDataLoader struct {
	getSingleRawBook RawBook
	getRawBooks      []RawBook

	getSingleRawDvd RawDvd
	getRawDvds      []RawDvd

	getSingleRawCd RawCd
	getRawCds      []RawCd
}

func NewRawDataLoader() RawDataLoader {
	singleRawBook := func() RawBook {
		path := "../json/book_single.json"
		return mustGetDataFromJson[RawBook](handler.ReadFileAsString(path))
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

	rawCd := func() RawCd {
		path := "../json/cd_single.json"
		return mustGetDataFromJson[RawCd](handler.ReadFileAsString(path))
	}

	rawCds := func() []RawCd {
		path := "../json/cd_arr.json"
		return mustGetDataFromJson[[]RawCd](handler.ReadFileAsString(path))
	}

	return RawDataLoader{
		getSingleRawBook: singleRawBook(),
		getRawBooks:      rawBooks(),
		getSingleRawDvd:  singleRawDvd(),
		getRawDvds:       rawDvds(),
		getSingleRawCd:   rawCd(),
		getRawCds:        rawCds(),
	}
}

func mustGetDataFromJson[T RawData](content string) T {
	var entity T
	err := json.Unmarshal([]byte(content), &entity)
	if err != nil {
		panic(err)
	}
	return entity
}
