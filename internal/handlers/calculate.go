package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Sanjay-R/gocalculator/api"
)

//custom unauthorized error to be used in this file
var ErrZeroInDenominator = errors.New("can't divide by 0")

//private function?
func getBody(w http.ResponseWriter, r *http.Request) (api.CalculatorRequest, error) {
	var body api.CalculatorRequest
	err := json.NewDecoder(r.Body).Decode(&body) // https://medium.com/what-i-talk-about-when-i-talk-about-technology/go-code-snippet-json-encoder-and-json-decoder-818f81864614

	if err != nil {
		slog.Error(err.Error())
		api.RequestErrorHandler(w, err)
		return body, err
	}
	return body, nil
}

//can do generics here: https://arc.net/l/quote/euranprl
func writeResponse[T api.CalculatorResponse](w http.ResponseWriter, response T) {
	//writte it to the response writer
	w.Header().Set("Content-Type", "application/json")
	err:= json.NewEncoder(w).Encode(response)
	if err != nil {
		slog.Error(err.Error())
		api.InternalErrorHandler(w)
		return
	}
}

func Addition(w http.ResponseWriter, r *http.Request) {
	// Assume call is already authenticated
	body, err := getBody(w, r)
	slog.Info("Request received for Addition", "body", body)

	if err != nil {
		return
	}

	response := api.CalculatorResponse {
		Answer: body.Left + body.Right,
	}

	writeResponse(w, response)
}

func Subtraction(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(w, r)
	slog.Info("Request received for Subtraction", "body", body)

	if err != nil {
		return
	}

	response := api.CalculatorResponse {
		Answer: body.Left - body.Right,
	}

	writeResponse(w, response)
}

func Multiplication(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(w, r)
	slog.Info("Request received for Multiplication", "body", body)

	if err != nil {
		return
	}

	response := api.CalculatorResponse {
		Answer: body.Left * body.Right,
	}

	writeResponse(w, response)
}

func Division(w http.ResponseWriter, r *http.Request) {
	body, err := getBody(w, r)

	if err != nil {
		return
	}

	if (body.Right == 0) {
		slog.Error("Invalid value 0 in denominator")
		api.RequestErrorHandler(w, ErrZeroInDenominator)
		return
	}

	slog.Info("Request received for Division", "body", body)
	response := api.CalculatorResponse {
		Answer: body.Left / body.Right,
	}

	writeResponse(w, response)
}
