package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteSuccess(w http.ResponseWriter, message string, data interface{}) {
	writeJSON(w, http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, message string, statusCode int) {
	writeJSON(w, statusCode, Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
