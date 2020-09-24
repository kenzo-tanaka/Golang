/*

- ref: https://hodalog.com/creating-restfulapi-with-golang/
- 記事一覧や詳細などを返すエンドポイントを実装
- DB接続などはやっていない

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome")
	fmt.Println("Endpoint Hit: Homepage")
}

// http://localhost:8081
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{}

	for i := 0; i < 10; i++ {
		title := "Hello_%d"
		articles = append(
			articles,
			Article{Title: fmt.Sprintf(title, i), Desc: "Article Desc", Content: "Article Content"})
	}
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key+"\n")
}

func main() {
	handleRequests()
}
