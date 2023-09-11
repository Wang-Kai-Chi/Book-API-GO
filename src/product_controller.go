package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
}

func (controller ProductController) QueryWithLimit(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectDB()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err == nil {
		var service ProductService
		json.NewEncoder(w).Encode(service.QueryWithLimit(db, 400))
	} else {
		panic(err)
	}
}

func GetProductsFromRequestBody(r *http.Request) ([]Product, error) {
	body, _ := io.ReadAll(r.Body)
	var ps []Product

	err := json.Unmarshal(body, &ps)

	return ps, err
}

func (controller ProductController) Insert(w http.ResponseWriter, r *http.Request) {
	ps, err := GetProductsFromRequestBody(r)

	if err == nil {
		db, err := ConnectDB()
		if err == nil {
			var service ProductService

			service.Insert(db, ps)
			json.NewEncoder(w).Encode(ps)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

}
func (controller ProductController) QueryWithPriceRange(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectDB()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err == nil {
		var ps ProductService
		min, err := strconv.Atoi(mux.Vars(r)["min"])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(mux.Vars(r)["max"])
		if err != nil {
			panic(err)
		}
		products := ps.QueryWithPriceRange(db, min, max)
		json.NewEncoder(w).Encode(products)
	} else {
		panic(err)
	}
}
