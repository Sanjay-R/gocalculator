package handlers

import (
	"log/slog"

	"github.com/Sanjay-R/gocalculator/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

//if func starts with capital letter -> tells the compiler that the function can be imported in other packages
//if starts with lower case -> private function, which could only be used in this handlers package
func ApiHandler(r *chi.Mux) {
	// Global middleware -> is applied all the time -> to all endpoints i make
	r.Use(chimiddle.StripSlashes) //will make sure trailing slashes are ignored. prevents 404 errors

	slog.Info("ApiHandler ready!")

	//takes in the path, as well as an anonymous function (which takes in a chi.Router as a param)
	r.Route("/add", func(router chi.Router) {

		slog.Debug("Route: {}", "/add", nil)

		// Middleware for /add route 
		// to check if user is authorized to access this data first
		router.Use(middleware.Authorization)

		// router.Post("/coins", GetCoinBalance) todo: add post req
	})
}