package api

import (
	"encoding/json"
	"net/http"
)

type CalculatorRequest struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type CalculatorResponse struct {
	Result int `json:"result"`
}

type DivisionRequest struct {
	Dividend int `json:"dividend"`
	Divisor int `json:"divisor"`
}

type SumRequest struct {
	Items []int `json:"items"`
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