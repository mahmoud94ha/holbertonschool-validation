package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"

  "github.com/gorilla/mux"
)

func main() {
  httpAddr := "0.0.0.0:9999"
  if port := os.Getenv("PORT"); port != "" {
    httpAddr = "0.0.0.0:" + port
  }
  fmt.Println("HTTP Server listening on", httpAddr)
  log.Fatal(http.ListenAndServe(httpAddr, setupRouter()))
}

func setupRouter() *mux.Router {
  r := mux.NewRouter()

  r.HandleFunc("/health", HealthCheckHandler).Methods("GET")

  r.HandleFunc("/hello", HelloHandler).Methods("GET")
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))

  return r
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("HIT: healthcheck")

  _, _ = io.WriteString(w, "ALIVE")

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  queryParams := r.URL.Query()
  nameParams := queryParams["name"]

  var name string
  switch len(nameParams) {
     case 0:
       name = "there"
     default:
       name = nameParams[0]
  }

  if name == "" {
    w.WriteHeader(400)
    return
  }

  _, _ = io.WriteString(w, fmt.Sprintf("Hello %s!", name))

  fmt.Printf("HIT: hello handler with name %s \n", name)
}
