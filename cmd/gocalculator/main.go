package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Sanjay-R/gocalculator/internal/handlers" //import package from my own module
	"github.com/go-chi/chi"                              //flexible and easy to use package for web dev
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger) //so anywhere I do slog.info it uses the logger just created

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	var r *chi.Mux = chi.NewRouter()
	handlers.ApiHandler(r)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%v", port)
	fmt.Println("Starting the GO stateless Calculator backend service at " + addr)

	err := http.ListenAndServe(addr, r) //start server with http package
	//takes the base location, in this case localhost:3000

	if err != nil {
		slog.Error(err.Error())
	}
}
