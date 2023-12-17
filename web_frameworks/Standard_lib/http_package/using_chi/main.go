package main

import (
  "net/http"

  "fmt"
  "log"
  "encoding/json"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)

type Article struct {
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Println()
  fmt.Fprintf(w, "Welcome to the Homepage!")
  fmt.Println("Endpoint Hit: homepage...")
  fmt.Println()
}

func helpPage(w http.ResponseWriter, r *http.Request) {
  fmt.Println()
  fmt.Fprintf(w, "Welcome to the Help page...")
  fmt.Println("Endpoint Hit: Help page...")
  fmt.Println()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
  fmt.Println()
  fmt.Println("Endpoint Hit: returnAllArticles...")
  json.NewEncoder(w).Encode(Articles)
  fmt.Println()
}

func getArticleByTitle(w http.ResponseWriter, r *http.Request) {
  fmt.Println()
  providedArticleTitle := chi.URLParam(r, "title")
  fmt.Println("Endpoint Hit: getArticleByTitle")
  for _, article := range Articles {
    if providedArticleTitle == article.Title {
      json.NewEncoder(w).Encode(article)
      return
    }
  }
  w.WriteHeader(404)
  w.Write([]byte("Error: article not found"))
  fmt.Println()
}

func returnNewRouter() *chi.Mux {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/", homePage)
  r.Get("/help", helpPage)
  r.Get("/articles", returnAllArticles)
  r.Get("/articles/{title}", getArticleByTitle)
  return r
}

func handleRequests() {
  r := returnNewRouter()

  fmt.Println("Chi Server Listening...")
  err := http.ListenAndServe(":10000", r)
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
