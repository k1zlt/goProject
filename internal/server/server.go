package server

import (
	"first/internal/api"
	"first/internal/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// func Run(address string, router http.Handler) {
func Run() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/add", middleware.Middleware(http.HandlerFunc(api.Add)))
	r.HandleFunc("/sub", middleware.Middleware(http.HandlerFunc(api.Subtraction)))
	r.HandleFunc("/mul", middleware.Middleware(http.HandlerFunc(api.Multiplication)))
	r.HandleFunc("/div", middleware.Middleware(http.HandlerFunc(api.Division)))
	fmt.Printf("The server is started!!!\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
