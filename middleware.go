package main

import (
	"log"
	"net/http"
	"time"
)

func myLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println()
		log.Println("***********************************************")
		log.Println(time.Now().Format(time.RFC850))
		log.Println(r.Method, r.URL)
		log.Println("***********************************************")
		defer log.Println()
		defer log.Println()

		next.ServeHTTP(w, r)
	})
}
