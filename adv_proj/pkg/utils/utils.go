package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func ParseBody (r *http.Request, X any) {
  if body, err := ioutil.ReadAll(r.Body); err == nil {
    if err := json.Unmarshal(body, X); err != nil {
      return
    }
  }
}
