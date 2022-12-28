package main

import (
	"net/http"

	"github.com/gorilla/mux"

  "api/pkg/routes"
)

func HomeHandler(res http.ResponseWriter, req *http.Request){}

func main() {
  r := mux.NewRouter()
  routes.RegisterBookStoreRoutes(r)
}
