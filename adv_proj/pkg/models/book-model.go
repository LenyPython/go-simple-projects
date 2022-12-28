package models

import(
  "github.com/jinzhu/gorm"
  "api/pkg/conf"
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
