package handlers

import "CRUD/strucs"


var Dir1 = strucs.Director{FirstName: "John", LastName: "Doe"}
var Dir2 = strucs.Director{FirstName: "Bob", LastName: "Dylan"}
var Dir3 = strucs.Director{FirstName: "Alan", LastName: "Diros"}

var moviesDB []strucs.Movie = []strucs.Movie{
  {Id: "1", Title: "First", Director: &Dir1},
  {Id: "2", Title: "Second", Director: &Dir2},
  {Id: "3", Title: "Third", Director: &Dir3},
}

