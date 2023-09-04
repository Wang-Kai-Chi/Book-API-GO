package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebController struct{}

func (w WebController) Init() {
	router := mux.NewRouter()

	var p ProductController
	router.HandleFunc("/product/query", p.QueryWithLimit).Methods("GET")
	router.HandleFunc("/product/insert", p.Insert).Methods("POST")
	router.HandleFunc("/product/query/min={min},max={max}", p.QueryWithPriceRange).Methods("GET")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
