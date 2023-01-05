package config

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
  DB *gorm.DB
)

func Connect(){
  db, err := gorm.Open(mysql.Open(DB_DNS), &gorm.Config{})
  if err != nil { panic(err) }
  DB = db
} 

func GetDB() *gorm.DB {
  return DB
}
