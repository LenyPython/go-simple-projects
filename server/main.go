package main

import(
  "fmt"
  "log"
  "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/hello" {
    http.Error(w, "404 not found", http.StatusNotFound)
    return
  }
  if r.Method != "GET" {
    http.Error(w, "Not supported", http.StatusNotFound)
    return
  }
  fmt.Println("Reached login")

}

func main() {
  fileServer := http.FileServer(http.Dir("./html"))
  http.Handle("/", fileServer)
  http.HandleFunc("/hello", helloHandler)

  fmt.Println("Starting on port 8000")
  if err := http.ListenAndServe(":8000", nil); err != nil {
    log.Fatal(err)
  }
}
