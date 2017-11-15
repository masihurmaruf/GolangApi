package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

type Article struct {
	Id 		int	   `json:id`
	Title 	string `json:"title"`
	Desc 	string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM articles")
	if err != nil {
		log.Print(err.Error())
	}
	var articles []Article

	for results.Next() {
		var article Article
		// for each row, scan the result into our tag composite object
		err = results.Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
			json.NewEncoder(w).Encode(HttpResp{Status: 404, Description: "Failed to select all from articles", Body: "Error"})
		}
		// and then print out the tag's Name attribute
		articles = append(articles, article)
	}
	fmt.Println("EndPoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	vars := mux.Vars(r)
	articleId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Print(err.Error())
	}
	var article Article
	err = db.QueryRow("SELECT * FROM articles where id = ?", articleId).Scan(&article.Id, &article.Title, &article.Desc, &article.Content)
	if err != nil {
		log.Print(err.Error())
		json.NewEncoder(w).Encode(HttpResp{Status: 404, Description: "Failed to select the article", Body: "Error"})
	} else {
		fmt.Println("EndPoint Hit: returnSingleArticle")
		json.NewEncoder(w).Encode(article)
	}
}

func InsertArticle(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var article Article
	err := decoder.Decode(&article)

	if err != nil {
		log.Print(err.Error())
	}

	stmt, _ := db.Prepare("INSERT INTO articles(title, description, content) VALUES (?,?,?)")
	res, err := stmt.Exec(article.Title, article.Desc, article.Content)

	if err != nil {
		log.Print(err.Error()) // proper error handling instead of panic in your app
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to insert article into database", Body: "ERROR"})
	}

	id, err := res.LastInsertId()
	if err != nil {
		json.NewEncoder(w).Encode(HttpResp{Status: 500, Description: "Failed to get last insert id", Body: "ERROR"})
	}
	fmt.Println("EndPoint Hit: InsertArticle")
	json.NewEncoder(w).Encode(HttpResp{Status: 200, Description: "Successfully Inserted Post Into the Database", Body: strconv.Itoa(int(id))})
}
