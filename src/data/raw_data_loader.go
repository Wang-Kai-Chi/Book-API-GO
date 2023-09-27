package data

import (
	"encoding/json"

	"iknowbook.com/handler"
)

func LoadData[T RawData](path string) T {
	mustGetDataFromJson := func(content string) T {
		var entity T
		err := json.Unmarshal([]byte(content), &entity)
		if err != nil {
			panic(err)
		}
		return entity
	}
	return mustGetDataFromJson(handler.ReadFileAsString(path))
}
