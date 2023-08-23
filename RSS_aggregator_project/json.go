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
