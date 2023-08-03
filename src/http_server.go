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
)

const (
	HELLO = "/hello"
	BOOKS = "/books"
)

func ServerStart() {
	fmt.Printf("Starting server at port 8080\n")

	http.Handle("/", http.FileServer(http.Dir("./static")))

	GetRequest(HELLO)
	GetRequest(BOOKS)

	PostRequest("/userinfo")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetRequest(pattern string) {
	switch pattern {
	case HELLO:
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello!")
		})
	case BOOKS:
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			books := GetBooksFromMongo()

			result, err := json.Marshal(books)

			if err == nil {
				fmt.Fprintf(w, string(result))
			}
		})
	}
}

func GetBooksFromMongo() []Book {
	ctx := context.TODO()
	queryStr := bson.D{{}}

	coll := GetMongoCollection("book")
	cur, err := coll.Find(ctx, queryStr)

	var books []Book

	if err == nil {
		for cur.Next(context.Background()) {
			var book Book
			err := cur.Decode(&book)
			if err == nil {
				books = append(books, book)
			}
		}
	}

	return books
}

func GetMongoCollection(collectionName string) *mongo.Collection {
	const URI = "mongodb://localhost:27017/?timeoutMS=5000"
	applyUri := options.Client().ApplyURI(URI)
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, applyUri)

	var coll *mongo.Collection

	if err == nil {
		coll = client.Database("iknowbook").Collection(collectionName)
	} else {
		panic(err)
	}
	return coll
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
