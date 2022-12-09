package main

import (
	// "fmt"
	// "log"
	// "math/rand"
	// "net/http"
	//"github.com/gorilla/mux"
	"fmt"
  "CRUD/strucs"
)



func main(){
  m := strucs.Movie{Id:"1", Title: "title"}
  fmt.Print(m.Id, " ", m.Title, "\n")
  

}
