package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
  DB *gorm.DB
)

func Connect(){
  db, err := gorm.Open("mysql", DBuser+":"+
    DBPass+"@/"+DBname+
    "?charset=utf8&parseTime-True&loc-Local")
  if err != nil {
    panic(err)
  }
  DB = db
} 

func GetDB() *gorm.DB {
  return DB
}
