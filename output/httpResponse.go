package output

import (
  "fmt"
  "encoding/json"
  "net/http"
)

type Response struct {
    StatusCode  int       `json:"statusCode"`
    Headers     map[string]string  `json:"headers"`
    Message        string    `json:"message"`
    Error        string    `json:"error"`
}

type Exception struct {
    Message string `json:"message"`
}

var resp Response

func ExceptionMessage(w http.ResponseWriter, message string, statusCode int){
  generateOutput(w, message, statusCode)
}

func HttpResponse(w http.ResponseWriter, message string, statusCode int){
  generateOutput(w, message, statusCode)
}

func generateOutput(w http.ResponseWriter, message string, statusCode int){
  resp.Message=fmt.Sprintf(message)
  resp.StatusCode=statusCode
  m := make(map[string]Response)
  m["data"] = resp
  json.NewEncoder(w).Encode(m)
}
