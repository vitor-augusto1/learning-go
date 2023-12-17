package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

type Article struct {
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to the Homepage!")
  fmt.Println("Endpoint Hit: homepage...")
}

func helpPage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to the Help page...")
  fmt.Println("Endpoint Hit: Help page...")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: returnAllArticles...")
  json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/help", helpPage)
  http.HandleFunc("/articles", returnAllArticles)

  fmt.Println("Server Listening...")
  err := http.ListenAndServe(":10000", nil)
  if err != nil {
    log.Fatal("Listen and Serve error. Cannot set up server...")
  }
}

func main() {
  Articles = []Article {
    Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    Article{Title: "Hello2", Desc: "Article Description", Content: "Article Content"},
  }
  handleRequests()
}
