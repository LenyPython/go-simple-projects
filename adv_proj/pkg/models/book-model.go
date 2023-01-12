package models

import (
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

func GetAllBooks() []Book {
  var books []Book
  result := db.Find(&books)
  if result.Error != nil { panic(result.Error) }
  return books
}

func GetABook(id uint64) (*Book, error) {
  var book Book
  result := db.Find(&book, id)
  if result.Error != nil || result.RowsAffected == 0 {
    return &book, errors.New("Error while gettin a book with id" + strconv.FormatUint(id, 10))
  }
  return &book, nil
}

func DeleteBook(id uint64) (*Book, error) {
  var book Book
  result := db.First(&book, id)
  if result.Error != nil {
    return &book, errors.New("Could not delete: " + strconv.FormatUint(id, 10))
  }
  result = db.Delete(&book)
  if result.Error != nil{
    return &book, errors.New("Error occured: "+result.Error.Error())
  }
  return &book, nil
}

func UpdateBook(ubook *Book) *Book {
  var prev Book
  db.First(&prev, ubook.ID)
  if ubook.Name != "" { prev.Name = ubook.Name }
  if ubook.Author != "" { prev.Author = ubook.Author }
  if ubook.Publication != "" { prev.Publication = ubook.Author }
  db.Save(&prev)
  return &prev
}
