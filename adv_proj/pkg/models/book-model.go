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

func DeleteBook(id uint64) Book {
  var book Book
  db.Where("ID=?",id).Delete(book)
  return book
}

func UpdateBook(ubook *Book) *Book {
  var prev Book
  db.Where("ID = ?", ubook.ID).First(&prev)
  if ubook.Name != "" { prev.Name = ubook.Name }
  if ubook.Author != "" { prev.Author = ubook.Author }
  if ubook.Publication != "" { prev.Publication = ubook.Author }
  db.Save(&prev)
  return &prev
}
