package config

var DBuser = "user"
var DBPass = "password"
var PROTOCOL = "tcp"
var ADDR = "(127.0.0.1:3306)"
var DBname = "go_db"
var APIport = "localhost:8001"

var DB_DNS = DBuser+":"+DBPass+"@"+PROTOCOL+ADDR+"/"+DBname+
"?charset=utf8mb4&parseTime=True&loc=Local"
