package main

import (
  "fmt"
  "encoding/json"
)

type User struct {
  Name string
}

func main() {
  user := &User{Name: "Frank"}
  b, err := json.Marshal(user)
  if err!= nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(b))
}
