package main

import (
	"auth-service/config"
	"auth-service/handlers"
	"net/http"
)

func routes(app *config.Config) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create-user", handlers.CreateHandlerFunc(app))
	mux.HandleFunc("/login", handlers.LoginHandlerFunc(app))

	return mux
}
