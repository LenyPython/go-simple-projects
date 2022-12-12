package handlers

import (
	"encoding/json"
	"net/http"
	"CRUD/strucs"
	"github.com/gorilla/mux"
)


func GetMovies(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-type", "application/json")
  json.NewEncoder(res).Encode(moviesDB)
}
func GetMovie(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-type", "application/json")
  var movie strucs.Movie
  params := mux.Vars(req)
  for _, item := range moviesDB {
    if item.Id == params["id"] {
      movie = item
      break
    }
  }
  json.NewEncoder(res).Encode(movie)
}
func DeleteMovie(res http.ResponseWriter, req *http.Request){
  res.Header().Set("Content-type","application/json")
  params := mux.Vars(req)
  for i, item := range moviesDB {
    if item.Id == params["id"]{
      moviesDB = append(moviesDB[:i],moviesDB[i+1:]...)
      break
    }
  }
}
