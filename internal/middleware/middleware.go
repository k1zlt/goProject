package middleware

import (
	"fmt"
	"net/http"
)

func Middleware(f http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request is made.")
		f.ServeHTTP(w, r)
	})
}
