package strucs

//import "encoding/json"


type Movie struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}
 type Director struct {
  FirstName string `json:"FirstName"`
  LastName string `json:"LastName"`
}
type UpdateData struct {
  Title string `json:"title"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}


