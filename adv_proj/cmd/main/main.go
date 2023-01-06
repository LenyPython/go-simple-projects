package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"api/pkg/config"
	"api/pkg/routes"
)


func main() {
  r := mux.NewRouter()
  routes.RegisterBookStoreRoutes(r)
  fmt.Println("Starting to Listen: ", config.APIport)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(config.APIport, r))
}
