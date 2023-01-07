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
  book := &models.Book{}
  utils.ParseBody(req, book)
  res.Header().Set("Content-Type", "application/json")
  b, err := models.CreateBook(book)
  var response json.RawMessage
  if err != nil {
    res.WriteHeader(http.StatusConflict)
    response, _ = json.Marshal(err)
  } else {
    res.WriteHeader(http.StatusOK)
    response, _ = json.Marshal(b)
  }
  res.Write(response)
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
  bookRes := models.GetABook(id)
  var jsonBook json.RawMessage
  res.Header().Set("Content-Type","application/json")
  if bookRes.ID == 0 { 
    jsonBook, _ = json.Marshal(map[string]string{"error":"No book found"})
    res.WriteHeader(http.StatusNotFound)
  } else { 
    jsonBook, _ = json.Marshal(bookRes)
    res.WriteHeader(http.StatusOK)
  }
  res.Write(jsonBook)
}

func DeleteBook(res http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  id, err := strconv.ParseUint(params["id"], 10, 64)
  if err != nil { panic(err) }
  bookDel := models.DeleteBook(id)
  jsonDel, _ := json.Marshal(bookDel)
  res.Header().Set("Content-Type","application/json")
  res.WriteHeader(http.StatusOK)
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
