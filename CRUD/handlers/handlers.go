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
  var data strucs.UpdateData
  movie := strucs.Movie{Id:"",Title:"", Director: &strucs.Director{FirstName:"",LastName:""}}
  json.NewDecoder(req.Body).Decode(&data)
  for i, item := range moviesDB {
    if item.Id == id {
      movie.Id = id
      movie.Director.FirstName = item.Director.FirstName
      movie.Director.LastName = item.Director.LastName
      movie.Title = item.Title
      moviesDB = append(moviesDB[:i], moviesDB[i+1:]...)
      break
    }
    if i == len(moviesDB) -1 {
      json.NewEncoder(res).Encode("Wrong id")
      return
    }
  }
  if data.Title != "" {
    movie.Title = data.Title
  }
  if data.FirstName != "" {
    movie.Director.FirstName = data.FirstName
  }
  if data.LastName != "" {
    movie.Director.LastName = data.LastName
  }
  moviesDB = append(moviesDB, movie)
  json.NewEncoder(res).Encode(movie)
}
