package utils

import (
	"encoding/json"
	"net/http"
)

type FailureBody struct {
	Status int
	ErrorMessage string
	Message string
}

type SuccessBody struct {
	Version string `json:"version"`
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


func Success(res http.ResponseWriter, data *SuccessBody) {
	res.WriteHeader(http.StatusAccepted);
	jsonData, _ := json.Marshal(data);
	res.Write(jsonData)
}

func Failure(res http.ResponseWriter, status int, data *FailureBody) {
	res.WriteHeader(status)
	jsonData, _ := json.Marshal(data);
	res.Write(jsonData)
}