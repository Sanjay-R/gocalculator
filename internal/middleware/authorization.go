package middleware

import (
	"log/slog"
	"net/http"

	"github.com/Sanjay-R/gocalculator/api"
	customError "github.com/Sanjay-R/gocalculator/internal/errors"
)

// auth function used as middleare
// thus gotta follow a certain signature -> take in- and return- http handler interface
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Authorizing")

		//define all logic for authorizing the http request
		var token = r.Header.Get("Authorization")

		if token != "devtoken" {
			slog.Error(customError.ErrUnAuthorized.Error())
			api.RequestErrorHandler(w, customError.ErrUnAuthorized)
			return
		}

		next.ServeHTTP(w, r) //in order to call the next middleware in line
		//or the Handler function for the endpoint if there's no middleware left
		//Authorization -> next.ServeHTTP -> Sum/Addition/Subtraction/Multiplication/Division
	})
}
