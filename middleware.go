package main

import (
	"fmt"
	"net/http"
	"time"
)

func myLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		fmt.Println("***********************************************")
		fmt.Println(time.Now().Format(time.RFC850))
		fmt.Println(r.Method, r.URL)
		fmt.Println("***********************************************")
		defer fmt.Println()
		defer fmt.Println()

		next.ServeHTTP(w, r)
	})
}
