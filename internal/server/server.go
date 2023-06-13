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
	r.Handle("/add", middleware.Middleware(api.Add))
	r.Handle("/sub", middleware.Middleware(api.Subtraction))
	r.Handle("/mul", middleware.Middleware(api.Multiplication))
	r.Handle("/div", middleware.Middleware(api.Division))
	fmt.Printf("The server is started!!!\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
