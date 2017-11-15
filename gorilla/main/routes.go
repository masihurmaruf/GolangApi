package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes =  Routes{
	Route{"HomePage", "GET", "/", HomePage},
	Route{"ReturnAllArticles", "GET", "/articles", ReturnAllArticles},
	Route{"ReturnSingleArticles", "GET", "/articles/{id}", ReturnSingleArticle},
	Route{"InsertArticle", "POST", "/articles", InsertArticle},
}
