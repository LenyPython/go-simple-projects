package service

import (
	"encoding/json"
	"net/http"

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
