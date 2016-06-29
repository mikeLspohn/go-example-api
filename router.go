package main

import (
	"net/http"
	// "os"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type Router struct {
	Port string
}

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
