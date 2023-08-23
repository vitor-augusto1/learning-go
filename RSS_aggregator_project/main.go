package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
  "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
  godotenv.Load(".env")

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT is not found in the environment.")
  }

}
