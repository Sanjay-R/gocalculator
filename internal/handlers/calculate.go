package handlers

import (
	// "encoding/json"
	"log/slog"
	"net/http"

	"github.com/Sanjay-R/gocalculator/api"
	"github.com/gorilla/schema"
)

func Addition(w http.ResponseWriter, r *http.Request) {
	// Assume call is already authenticated
	var body = api.CalculatorRequest{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error = decoder.Decode(&body, r.Form) //todo: check how to decode body

	if(err != nil) {
		slog.Error(err.Error())
		api.InternalErrorHandler(w)
		return
	}
}