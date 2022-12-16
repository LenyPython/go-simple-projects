package handlers

import (
	"CRUD/strucs"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

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
  params := mux.Vars(req)
  id := params["id"]
  title := params["title"]
  director := strings.Split(params["dir"], " ")
  for _, movie := range moviesDB {
    if movie.Id == id {
      movie.Title = title
      movie.Director = &strucs.Director{FirstName:director[0], LastName: director[1]}
      break
    }
  }
}
