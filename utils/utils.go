package utils

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"time"
)

func GenerateRandomNumber()int{
rand.Seed(time.Now().UnixNano()) 
return rand.Intn(900000) + 100000

}

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