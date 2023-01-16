package routes

import (
  "github.com/gorilla/mux"

  "api/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *mux.Router){
  r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
  r.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
  r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
  r.HandleFunc("/book/{id}", controllers.UpdateBookWithId).Methods("PUT")
  r.HandleFunc("/book/{id}", controllers.DeleteBookById).Methods("DELETE")
}
