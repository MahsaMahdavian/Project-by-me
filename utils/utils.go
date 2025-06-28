package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter,statusCode int,payload interface{}){

 response,err:=json.Marshal(payload)
 if err!=nil{
	http.Error(w,err.Error(),http.StatusInternalServerError)
	return
 }

 w.Header().Set("Content_Type","application/json")
 w.WriteHeader(statusCode)
 w.Write(response)
}

func ResponseWithError(w http.ResponseWriter,statusCode int,message string){
	ResponseWithJson(w,statusCode,map[string]string{"error":message})
}