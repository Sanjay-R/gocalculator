package middleware

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Sanjay-R/gocalculator/api"
)

//custom unauthorized error to be used in this file
var ErrUnAuthorized = errors.New("invalid username or token")

// auth function used as middleare
// thus gotta follow a certain signature -> take in- and return- http handler interface
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Authorizing")
		
		//define all logic for authorizing the http request
		var token = r.Header.Get("Authorization")

		if (token != "devtoken") {
			slog.Error(ErrUnAuthorized.Error())
			api.RequestErrorHandler(w, ErrUnAuthorized)
			return
		}

		next.ServeHTTP(w, r) //in order to call the next middleware in line
		//or the Handler function for the endpoint if there's no middleware left
		//Authorization -> next.ServeHTTP -> GetCoinBalance
	})
}