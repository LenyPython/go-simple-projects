package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(res http.ResponseWriter, req *http.Request){}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
}
