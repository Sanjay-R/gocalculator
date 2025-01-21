package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Sanjay-R/gocalculator/api"
)

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

func Addition(w http.ResponseWriter, r *http.Request) {
	// Assume call is already authenticated
	body, err := getBody(w, r)

	if err != nil {
		return
	}

	slog.Info("Request received", "body", body)
	response := api.CalculatorResponse {
		Answer: body.Left + body.Right,
	}

	//writte it to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		slog.Error(err.Error())
		api.InternalErrorHandler(w)
		return
	}
}
