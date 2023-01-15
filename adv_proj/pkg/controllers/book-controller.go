package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
  id, err := strconv.ParseUint(params["id"], 10, 64)
  if err != nil { panic(err) }
  bookDel, err := models.DeleteBook(id)
  res.Header().Set("Content-Type","application/json")
  var jsonDel json.RawMessage
  if err != nil {
    jsonDel, _ = json.Marshal(map[string]string{"error":err.Error()})
    res.WriteHeader(http.StatusNotFound)
  } else {
    jsonDel, _ = json.Marshal(bookDel)
    res.WriteHeader(http.StatusOK)
  }
  res.Write(jsonDel)
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
