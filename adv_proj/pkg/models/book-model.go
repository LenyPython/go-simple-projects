package models

import (
	"encoding/json"
	"errors"
	"strconv"

	"gorm.io/gorm"

	"api/pkg/config"
)

var db * gorm.DB

type Book struct {
  gorm.Model
  Name string `json:"name"`
  Author string `json:"author"`
  Publication string `json:"publication"`
}

func init() {
  db = config.Connect()
  db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) (*Book, error) {
  result := &Book{}
  res := db.FirstOrCreate(result, b)
  if res.RowsAffected == 0 {
    return result, errors.New("Entry exists")
  }
  return result, nil
}

func GetAllBooks() ([]Book, error) {
  var books []Book
  result := db.Find(&books)
  if result.Error != nil { 
    return books, errors.New(result.Error.Error())
  }
  return books, nil
}

func GetBookById(id uint64) (*Book, error) {
  var book Book
  result := db.Find(&book, id)
  if result.Error != nil || result.RowsAffected == 0 {
    return &book, errors.New("Error while gettin a book with id: " + strconv.FormatUint(id, 10))
  }
  return &book, nil
}

func DeleteBook(id uint64) (*Book, error) {
  var book Book
  result := db.First(&book, id)
  if result.Error != nil {
    return &book, errors.New("Book with this Id not found: " + strconv.FormatUint(id, 10))
  }
  result = db.Delete(&book)
  if result.Error != nil {
    return &book, errors.New("Error occured: "+result.Error.Error())
  }
  return &book, nil
}

func UpdateBook(book *Book) (json.RawMessage, error) {
  result := db.Save(book)
  response, _ := json.Marshal(book)
  if result.Error != nil {
    return response, errors.New("SOmething went wrong...")
  }
  return response, nil
}
