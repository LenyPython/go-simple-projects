package models

import(
  "fmt"

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

func inti() {
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
  result := db.Create(&b)
  if result.Error != nil {
    fmt.Println("Error while creating a book")
  }
  return b
}

func GetAllBooks() []Book {
  var books []Book
  db.Find(&books)
  return books
}

func GetABook(id uint64) *Book {
  var book Book
  result := db.Find(&book, id)
  if result.Error != nil {
    fmt.Println("Error while getting a book with id: ", id)
  }
  return &book
}

func DeleteBook(id uint64) Book {
  var book Book
  result := db.Delete(&book, id)
  if result.Error != nil {
    fmt.Println("Could not delete: ", id)
  }
  return book
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
