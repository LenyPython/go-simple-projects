package handlers

import (
  "net/http"
  "CRUD/strucs"
)

func GetMovies(r http.ResponseWriter, w *http.Request) []strucs.Movie{
  arr := []strucs.Movie{
    {Id: "1",Title: "FIrst mov"},
    {Id: "2",Title: "Sec mov"},
  }
  return arr
}
func GetMovie(r http.ResponseWriter, w *http.Request) strucs.Movie{
  return strucs.Movie{Id:"1", Title: "First mov"}
}
