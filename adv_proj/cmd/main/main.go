package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"api/pkg/routes"
  "api/pkg/conf"
)


func main() {
  r := mux.NewRouter()
  routes.RegisterBookStoreRoutes(r)
  log.Fatal(http.ListenAndServe(config.APIport, r))
}
