package main

import "net/http"

func handlerReadiness(rw http.ResponseWriter, req *http.Request) {
  respondWithJSON(rw, 200, struct{}{})
}
