package models

import(
  "github.com/jinzhu/gorm"
  "api/pkg/config"
)

var db * gorm.DB

type Book struct {
  gorm.Model
  Name string `json:"name"`
  Author string `json:"author"`
  Publication string `json:"publication"`
}

func inti() {
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
  db.NewRecord(b)
  db.Create(&b)
  return b
}

func GetAllBooks() []Book {
  var books []Book
  db.Find(&books)
  return books
}

func GetABook(id uint64) (*Book, *gorm.DB) {
  var book Book
  db := db.Where("ID=?", id).Find(&book)
  return &book, db
}
