package controllers

import (
	"net/http"

	"api/pkg/models"
	"api/pkg/service"

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

func DeleteBookById(res http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  deletedBook, status := service.DeleteBookById(params["id"])
  res.Header().Set("Content-Type","application/json")
  res.WriteHeader(status)
  res.Write(deletedBook)
}

func UpdateBookWithId(res http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  msg, status := service.UpdateBook(params["id"], req)
  res.Header().Set("Content-Type","application/json")
  res.WriteHeader(status)
  res.Write(msg)
}
