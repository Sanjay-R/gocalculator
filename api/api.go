package api

import (
	"encoding/json"
	"net/http"
)

type CalculatorRequest struct {
	Left int32
	Right int32
}

type CalculatorResponse struct {
	Answer int32
}

//Error response when an error occurs
type Error struct {
	Code int
	Message string
}

//used to write an error response to the HTTP response writer
//will return error response to the person who called the endpoint
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error {
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

//to have different types of error responses -> make wrapper below
var (
	RequestErrorHandler = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)