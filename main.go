package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8082"
	} else {
		port = ":" + port
	}

	router := Router{Port: port}
	r := router.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})
	handler := c.Handler(r)

	log.Println("Now listening on port", port)
	log.Fatal(http.ListenAndServe(router.Port, handler))
}
