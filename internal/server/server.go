package server

import (
	"first/internal/middleware"
	"first/pkg/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// func Run(address string, router http.Handler) {
func Run() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/lesson", middleware.Middleware(http.HandlerFunc(handlers.GetLesson)))
	fmt.Printf("The server is started!!!\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
