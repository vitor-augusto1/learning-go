package main

import (
  "errors"
  "fmt"
)

type User struct {
  username string
  email string
}

func main() {
  _, err := getUser("", "")
  if err != nil {
    fmt.Println("There is an error here: ", err)
    return
  }
  // Use user here
}

func getUser(username string, email string) (User, error) {
  if username == "" || email == "" {
    return User{}, errors.New("Empty required field")
  }
  u := User {username, email}
  return u, nil
}
