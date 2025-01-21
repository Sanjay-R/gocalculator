package middleware

import (
	"errors"
	"net/http"
	log "log/slog"
	"github.com/Sanjay-R/gocalculator/api"
)

//custom unauthorized error to be used in this file
var UnAuthorizedError = errors.New("Invalid username or token.")

// auth function used as middleare
// thus gotta follow a certain signature -> take in- and return- http handler interface
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		//define all logic for authorizing the http request
		var token = r.Header.Get("Authorization")

		if (token == "" || token != "random") {
			log.Error(UnAuthorizedError.Error())
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r) //in order to call the next middleware in line
		//or the Handler function for the endpoint if there's no middleware left
		//Authorization -> next.ServeHTTP -> GetCoinBalance
	})
}