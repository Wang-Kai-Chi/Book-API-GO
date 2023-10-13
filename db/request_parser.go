package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	. "iknowbook.com/data"
)

func MustGetEntityFromRequest[T Data](r *http.Request) []T {
	body, err := io.ReadAll(r.Body)

	var ps []T
	if err == nil {
		if len(body) == 0 {
			fmt.Println("Empty body")
		}
		err := json.Unmarshal(body, &ps)
		if err != nil {
			fmt.Println("not a list")
		}
	} else {
		panic(err)
	}
	return ps
}

func GetEntityFromBody[T Data](body []byte) ([]T, error) {
	var ps []T
	if len(body) == 0 {
		fmt.Println("Empty body")
	}
	err := json.Unmarshal(body, &ps)

	return ps, err
}
