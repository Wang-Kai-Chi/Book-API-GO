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

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func (controller ProductController) QueryWithLimit(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectDB()
	setHeader(w)
	if err == nil {
		var service ProductService
		json.NewEncoder(w).Encode(service.QueryWithLimit(db, 400))
	} else {
		panic(err)
	}
}

func getProductsFromRequestBody(r *http.Request) ([]Product, error) {
	body, _ := io.ReadAll(r.Body)
	var ps []Product

	err := json.Unmarshal(body, &ps)

	return ps, err
}

func (controller ProductController) Insert(w http.ResponseWriter, r *http.Request) {
	ps, err := getProductsFromRequestBody(r)
	setHeader(w)
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
	setHeader(w)
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

func (controller ProductController) QueryById(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectDB()
	setHeader(w)
	if err == nil {
		var ps ProductService
		products := ps.QueryByBarcode(db, mux.Vars(r)["barcode"])
		json.NewEncoder(w).Encode(products)
	} else {
		panic(nil)
	}
}
