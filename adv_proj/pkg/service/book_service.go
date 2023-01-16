package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/pkg/models"
	"api/pkg/utils"
)

func CreateBook(r *http.Request) (json.RawMessage, int) {
  book := &models.Book{}
  utils.ParseBody(r, book)
  b, err := models.CreateBook(book)
  var response json.RawMessage
  if err != nil {
    response, _ = json.Marshal(map[string]string{"error":err.Error()})
    return response, http.StatusConflict
  } 
  response, _ = json.Marshal(b)
  return response, http.StatusOK
}

func GetAllBooks() (json.RawMessage, int){
  res, err := models.GetAllBooks()
  var response json.RawMessage
  if err != nil {
    response, _ = json.Marshal(map[string]string{"error":err.Error()})
    return response, http.StatusConflict
  } 
  response, _ = json.Marshal(res)
  return response, http.StatusOK
}
func GetBookById(id_str string) (json.RawMessage, int) {
  var response json.RawMessage
  id, err := strconv.ParseUint(id_str, 10, 32)
  if err != nil { 
    response, _ = json.Marshal(map[string]string{"error":err.Error()})  
    return response, http.StatusBadRequest
  }
  bookRes, err := models.GetBookById(id)
  if err != nil { 
    response, _ = json.Marshal(map[string]string{"error":err.Error()})
    return response, http.StatusNotFound
  } 
  response, _ = json.Marshal(bookRes)
  return response, http.StatusOK
}
func DeleteBookById(id_str string) (json.RawMessage, int) {
  var response json.RawMessage
  id, err := strconv.ParseUint(id_str, 10, 64)
  if err != nil {
    response, _ := json.Marshal(map[string]string{"error":err.Error()})
    return response, http.StatusBadRequest
  }
  bookDel, err := models.DeleteBook(id)
  if err != nil {
    response, _ = json.Marshal(map[string]string{"error":err.Error()})
    return response, http.StatusNotFound
  } 
  response, _ = json.Marshal(bookDel)
  return response, http.StatusOK
}
func UpdateBook(id_str string, r *http.Request) (json.RawMessage, int) {
  var book models.Book
  var response json.RawMessage
  utils.ParseBody(r, &book)
  id, err := strconv.ParseUint(id_str, 10, 64)
  if err != nil {
    response, _ = json.Marshal(map[string]string{"error":"Bad id param"}) 
    return response, http.StatusBadRequest
  }
  currentBook, err := models.GetBookById(id)
  if  err != nil {
    response, _ = json.Marshal(map[string]string{"error":"Could not update book: "+err.Error()})
    return response, http.StatusNotFound
  }
  if book.Name != "" { currentBook.Name = book.Name }
  if book.Author != "" { currentBook.Author = book.Author }
  if book.Publication != "" { currentBook.Publication = book.Author }
  updatedBook, err := models.UpdateBook(currentBook)
  if err != nil {
    response, _ = json.Marshal(map[string]string{"error":"Could not update book: "+err.Error()})
    return response, http.StatusNotModified
  }
  response, _ = json.Marshal(updatedBook)
  return response, http.StatusOK

}
