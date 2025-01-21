package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi" //flexible and easy to use package for web dev
    "github.com/Sanjay-R/gocalculator/internal/handlers" //import package from my own module
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    slog.SetDefault(logger) //so anywhere I do slog.info it uses the logger just created

    var r *chi.Mux = chi.NewRouter()
    handlers.ApiHandler(r)

    port := 3000
    addr := fmt.Sprintf(":%d", port)
    fmt.Println("Starting the GO stateless Calculator backend service at " + addr)

    err := http.ListenAndServe(addr, r) //start server with http package
	//takes the base location, in this case localhost:3000

    if(err != nil) {
        slog.Error(err.Error());
    }
}