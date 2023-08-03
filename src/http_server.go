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
	HELLO     = "/hello"
	BOOKS     = "/books"
	USER_INFO = "/userinfo"
)

func ServerStart() {
	fmt.Printf("Starting server at port 8080\n")

	http.Handle("/", http.FileServer(http.Dir("./static")))

	getRequest(HELLO)
	getRequest(BOOKS)

	postRequest(USER_INFO)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getRequest(pattern string) {
	switch pattern {
	case HELLO:
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello!")
		})
	case BOOKS:
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			books := getBooksFromMongo()

			result, err := json.Marshal(books)

			if err == nil {
				fmt.Fprintf(w, string(result))
			}
		})
	}
}

func getBooksFromMongo() []Book {
	ctx := context.TODO()
	queryStr := bson.D{{}}

	coll := getMongoCollection("book")
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

func getMongoCollection(collectionName string) *mongo.Collection {
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

func postRequest(pattern string) {
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
