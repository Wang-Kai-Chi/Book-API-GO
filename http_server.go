package main

import (
	"fmt"
	"log"
	"net/http"
)

func ServerStart() {
	fmt.Printf("Starting server at port 8080\n")

	http.Handle("/", GetFileServer("./static"))

	http.HandleFunc("/hello", GetHello)
	http.HandleFunc("/userinfo", PostUserInfo)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func GetFileServer(dirPath string) http.Handler {
	return http.FileServer(http.Dir(dirPath))
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func PostUserInfo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST UserInfo successful")

	name := r.FormValue("uname")
	email := r.FormValue("uemail")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}
