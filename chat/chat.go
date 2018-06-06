package main

import (
  "net/http"
  "log"
  "time"
  "fmt"
)

func main() {
  http.HandleFunc("/welcome", messageHandler)

  server := &http.Server{
    Addr:           ":8080",
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  log.Println("Listening..."")
  server.ListenAndServe()
}

func ListenAndServe(addr string, handler Handler) error {
  server := &Server{Addr: addr, Handler: handler}
  return server.ListenAndServe()
}
