package config

var DBuser = "admin"
var DBPass = "password"
var PROTOCOL = "tcp"
var ADDR = "(127.0.0.1:8001)"
var DBname = "database"
var APIport = "localhost:8001"

var DB_DNS = DBuser + ":" + DBPass +
"@" + PROTOCOL + ADDR + "/" + DBname +
"?charset=utf8mb4&parseTime=True&locLocal"
