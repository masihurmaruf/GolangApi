package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Path("/").HandlerFunc(HomePage)
	myRouter.Path("/articles").HandlerFunc(ReturnAllArticles)
	myRouter.Path("/articles/{id}").HandlerFunc(ReturnSingleArticle)
	myRouter.Path("/articles").HandlerFunc(InsertArticle)

}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc).
			Headers("Content-Type", "application/json")
	}
	return router
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":10000", router))
}