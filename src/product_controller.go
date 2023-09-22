package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
}

func (ctr ProductController) QueryWithLimit(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectDB()
	setHeader(w)
	if err == nil {
		var service ProductService
		json.NewEncoder(w).Encode(service.QueryWithLimit(db, 400))
	} else {
		panic(err)
	}
}

func mustGetProductsFromRequest(r *http.Request) []Product {
	body, err := io.ReadAll(r.Body)

	var ps []Product
	if err == nil {
		if len(body) == 0 {
			fmt.Println("Empty body")
		}
		err := json.Unmarshal(body, &ps)
		if err != nil {
			fmt.Println("not list of products")
		}
	} else {
		panic(err)
	}

	return ps
}

func (ctr ProductController) Insert(w http.ResponseWriter, r *http.Request) {
	ps := mustGetProductsFromRequest(r)
	setHeader(w)

	db, err := ConnectDB()
	if err == nil {
		var service ProductService

		service.Insert(db, ps)
		json.NewEncoder(w).Encode(ps)
	} else {
		panic(err)
	}
}
func (ctr ProductController) QueryWithPriceRange(w http.ResponseWriter, r *http.Request) {
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

func (ctr ProductController) QueryById(w http.ResponseWriter, r *http.Request) {
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

func (ctr ProductController) Update(w http.ResponseWriter, r *http.Request) {
	ps := mustGetProductsFromRequest(r)
	setHeader(w)
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	var service ProductService
	service.Update(db, ps)
	fmt.Fprintf(w, "update successfully")
}

func (ctr ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	ps := mustGetProductsFromRequest(r)
	setHeader(w)
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	var service ProductService
	service.Delete(db, ps)
	fmt.Fprintln(w, "delete successfully")
}
