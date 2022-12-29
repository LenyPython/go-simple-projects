package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/pkg/models"
	"api/pkg/utils"

	"github.com/gorilla/mux"
)

var NewVook models.Book

func CreateBook(res http.ResponseWriter, req *http.Request){

}
func GetAllBooks(res http.ResponseWriter, req *http.Request){
  allBooks := models.GetAllBooks()
  jsonBooks, _ := json.Marshal(allBooks)
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(http.StatusOK)
  res.Write(jsonBooks)
}

func GetBook(res http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  id, err := strconv.ParseUint(params["id"], 10, 64)
  if err != nil { panic(err) }
  bookRes, _ := models.GetABook(id)
  jsonBook, _ := json.Marshal(bookRes)
  res.Header().Set("Content-Type","application/json")
  res.WriteHeader(http.StatusOK)
  res.Write(jsonBook)
}
