package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(rw http.ResponseWriter, code int, payload interface{}) {
  dat, err := json.Marshal(payload)
  if err != nil {
    log.Printf("Failed to Marshal JSON response: %v", payload)
    rw.WriteHeader(500)
    return
  }
  rw.Header().Add("Content-Type", "application/json")
  rw.WriteHeader(code)
  rw.Write(dat)
}

func respondWithError(rw http.ResponseWriter, code int, msg string) {
  if code > 499 {
    log.Println("Responding with 5XX error: ", msg)
  }

  type errResponse struct {
    Error string `json:"error"`
  }

  respondWithJSON(rw, code, errResponse{
    Error: msg,
  })
}
