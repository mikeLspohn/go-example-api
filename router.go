package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

// Router models the router with it's port that the server will bind too
type Router struct {
	Port string
}

// NewRouter sets up a new router with the given routes and middleware
func (r *Router) NewRouter() http.Handler {
	router := mux.NewRouter()
	// db := DBConnection()

	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", CreateUser).Methods("POST")

	chain := alice.New(myLogger).Then(router)
	return chain
}
