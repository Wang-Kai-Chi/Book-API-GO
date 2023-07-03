package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"iknowbook.com/handler"
)

func ServerStart() {
	fmt.Printf("Starting server at port 8080\n")

	http.Handle("/", GetFileServer("./static"))

	GetRequest("/hello")
	GetRequest("/books")

	PostRequest("/userinfo")

	fmt.Println(GetMongoBookCollec())

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func GetFileServer(dirPath string) http.Handler {
	return http.FileServer(http.Dir(dirPath))
}

func GetRequest(pattern string) {
	const HELLO = "/hello"
	const BOOKS = "/books"

	if pattern == HELLO {
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello!")
		})
	}
	if pattern == BOOKS {
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, handler.ReadFileAsString("./json/book.json"))
		})
	}
}

func GetMongoBookCollec() string {
	uri := "mongodb://localhost:27017/?timeoutMS=5000"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	var result bson.M

	if err != nil {
		panic(err)
	}
	coll := client.Database("iknowbook").Collection("book")
	err = coll.FindOne(context.TODO(), bson.D{{"price", "90元"}}).Decode(&result)

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		panic(err)
	}

	return string(jsonData)
}

func PostRequest(pattern string) {
	const USER_INFO = "/userinfo"

	if pattern == USER_INFO {
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
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
		})
	}
}
