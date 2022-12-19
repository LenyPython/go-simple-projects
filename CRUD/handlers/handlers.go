package handlers

import (
	"CRUD/strucs"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

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
func SaveMovie(res http.ResponseWriter, req *http.Request){
  res.Header().Set("Content-type", "application/json")
  var movie strucs.Movie
  _ = json.NewDecoder(req.Body).Decode(&movie)
  movie.Id = strconv.Itoa(rand.Intn(100000))
  moviesDB = append(moviesDB, movie)
  json.NewEncoder(res).Encode(movie)
}
func DeleteMovie(res http.ResponseWriter, req *http.Request){
  res.Header().Set("Content-type","application/json")
  params := mux.Vars(req)
  var deleted strucs.Movie
  for i, item := range moviesDB {
    if item.Id == params["id"]{
      deleted = item
      moviesDB = append(moviesDB[:i],moviesDB[i+1:]...)
      break
    }
  }
  json.NewEncoder(res).Encode(deleted)
}
func UpdateMovie(res http.ResponseWriter, req *http.Request){
  res.Header().Set("Content-type", "application/json")
  params := mux.Vars(req)
  id := params["id"]
  title := params["title"]
  directorName := params["dir_name"]
  directorLastName := params["dir_last"]
  var dir strucs.Director
  for i, item := range moviesDB {
    if item.Id == id {
      dir = *item.Director
      moviesDB = append(moviesDB[:i],moviesDB[i+1:]...)
      break
    }
  }
  if directorName != "" {
    dir.FirstName = directorName
  }
  if directorLastName != "" {
    dir.LastName = directorLastName
  }
  moviesDB = append(moviesDB, strucs.Movie{Id: id, Title: title, Director: &dir})
}
