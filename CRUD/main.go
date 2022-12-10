package main

import (
	// "math/rand"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"CRUD/handlers"
	"CRUD/strucs"
)


var PORT = ":8000"

func main(){
  m := strucs.Movie{Id:"1", Title: "title"}
  fmt.Print(m.Id, " ", m.Title, "\n")
  r := mux.NewRouter()
  r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
  r.HandleFunc("/movie/{id}", handlers.GetMovie).Methods("GET")
  r.HandleFunc("/movie", handlers.SaveMovie).Methods("POST")
  r.HandleFunc("/movie/{id}", handlers.UpdateMovie).Methods("PUT")
  r.HandleFunc("/movie/{id}", handlers.DeleteMovie).Methods("DELETE")

  fmt.Println("Starting server at port", PORT)
  log.Fatal(http.ListenAndServe(PORT, r))
  

}
