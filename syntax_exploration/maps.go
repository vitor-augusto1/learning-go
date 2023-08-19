package main

import (
  "fmt"
  "errors"
)

type user struct {
  name string
  phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
  userMap := make(map[string]user)
  if len(names) != len(phoneNumbers) {
    return nil, errors.New("Arrays length isn't equal")
  }
  for i := 0; i < len(names); i++ {
    newUser := user {name: names[i], phoneNumber: phoneNumbers[i],}
    userMap[names[i]] = newUser
  }
  return userMap, nil
}

func checkIfNameIsInMap(userMap map[string]user, name string) (bool, error) {
  _, ok := userMap[name]
  if ok == false {
    return false, errors.New("Key error, invalid key for map")
  }
  return ok, nil
}

func main() {
  names := []string {"Vitor", "Augusto", "Andrade"}
  phoneNumbers := []int {31997221726, 71876426381, 6382173625}
  usersMap, err := getUserMap(names, phoneNumbers)
  if err != nil {
    fmt.Println("ERROR: ", err)
    return
  }
  fmt.Println(usersMap)
  fmt.Println("Checking if 'Vitor' exists in the map...")
  exists, err := checkIfNameIsInMap(usersMap, "Vitor")
  if err != nil {
    fmt.Println("ERROR: ", err)
    return
  }
  fmt.Println(exists)
}
