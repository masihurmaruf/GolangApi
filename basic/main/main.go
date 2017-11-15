// Basic REST api in go
package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
)

type Article struct {
	Title 	string `json:"Title"`
	Desc 	string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", 	Desc: "Article Description", 	Content: "Articel Content"},
		Article{Title: "Hello 2", 	Desc:"Articel Description 2", 	Content: "Article Content 2"},
	}
	fmt.Println("EndPoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
