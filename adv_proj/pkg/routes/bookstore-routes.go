package routes

import (
  "github.com/gorilla/mux"
  "api/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *mux.Router){
  r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
  r.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
  r.HandleFunc("/book/{id}", controllers.GetBook).Methods("GET")
  r.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
  r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
