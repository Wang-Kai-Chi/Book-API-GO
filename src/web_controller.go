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
	router.HandleFunc("/product/query/barcode={barcode}", p.QueryById).Methods("GET")
	router.HandleFunc("/product/update", p.Update).Methods("PUT")

	router.HandleFunc("/product/delete", p.Delete).Methods("DELETE")

	static := http.Dir("./src/static/web/")
	router.PathPrefix("/").Handler(http.FileServer(static))
	port := ":8080"
	println("server start at localhost" + port)
	err := http.ListenAndServe(port, router)

	if err != nil {
		panic(err)
	}
}
