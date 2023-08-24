package main

import "net/http"


func handlerErr(rw http.ResponseWriter, req *http.Request) {
  respondWithError(rw, 400, "Something went wrong")
}
