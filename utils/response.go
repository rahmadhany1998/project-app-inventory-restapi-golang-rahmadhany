package utils

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	Limit        int `json:"limit"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

type Response struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

func WriteSuccess(w http.ResponseWriter, message string, statusCode int, data interface{}, pagination *Pagination) {
	writeJSON(w, statusCode, Response{
		Status:     "success",
		Message:    message,
		Data:       data,
		Pagination: pagination,
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
