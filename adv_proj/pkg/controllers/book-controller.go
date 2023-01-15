package controllers

import (
	"encoding/json"
	"net/http"

	"api/pkg/models"
	"api/pkg/service"
	"api/pkg/utils"

	"github.com/gorilla/mux"
)

var NewVook models.Book

func CreateBook(res http.ResponseWriter, req *http.Request){
  msg, status := service.CreateBook(req)
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(status)
  res.Write(msg)
}

func GetAllBooks(res http.ResponseWriter, req *http.Request){
  msg, status := service.GetAllBooks()
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(status)
  res.Write(msg)
}

func GetBookById(res http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  book, status := service.GetBookById(params["id"])
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(status)
  res.Write(book)
}

func DeleteBook(res http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  deletedBook, status := service.DeleteBookById(params["id"])
  res.Header().Set("Content-Type","application/json")
  res.WriteHeader(status)
  res.Write(deletedBook)
}

func UpdateBook(res http.ResponseWriter, req *http.Request){
  var book models.Book
  utils.ParseBody(req, &book)
  updated := models.UpdateBook(&book)
  updatedJson, err := json.Marshal(*updated)
  if err != nil { panic(err) }
  res.Header().Set("Content-Type","application/json")
  res.WriteHeader(http.StatusOK)
  res.Write(updatedJson)
}
