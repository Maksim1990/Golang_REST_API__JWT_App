package output

import (
  "fmt"
  "encoding/json"
    "github.com/goRESTapi/models"
    "net/http"
)


var resp models.Response

func ExceptionMessage(w http.ResponseWriter, message string, statusCode int){
  generateOutput(w, message, statusCode)
}

//func JSONResponse(w http.ResponseWriter,user models.User){
func JSONResponse(w http.ResponseWriter,i interface{}){
    m := make(map[string]models.User)
    m["data"] = user

    json.NewEncoder(w).Encode(m)
}

func JSONListResponse(w http.ResponseWriter,res []models.User){
    m := make(map[string][]models.User)
    m["data"] = res

    json.NewEncoder(w).Encode(m)
}

func HttpResponse(w http.ResponseWriter, message string, statusCode int){
  generateOutput(w, message, statusCode)
}

func generateOutput(w http.ResponseWriter, message string, statusCode int){
  resp.Message=fmt.Sprintf(message)
  resp.StatusCode=statusCode
  m := make(map[string]models.Response)
  m["data"] = resp
  json.NewEncoder(w).Encode(m)
}
