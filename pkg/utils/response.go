package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode    interface{} `json:"statusCode"`
	StatusMessage interface{} `json:"statusMessage"`
	Error         interface{} `json:"errorMessage"`
	Data          interface{} `json:"data"`
}

func RespondWithJSON(w http.ResponseWriter, statusMessage, statusCode, err, data interface{}) bool {
	isError := false
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resp := &Response{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		Error:         err,
		Data:          data,
	}

	r, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorJson, _ := json.Marshal(err.(error).Error())
		_, err := w.Write(errorJson)
		if err != nil {
			isError = true
		}
	}
	_, err2 := w.Write(r)
	return isError || err2 != nil
}
