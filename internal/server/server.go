package server

import (
	"first/internal/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// func Run(address string, router http.Handler) {
func Run() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/add", api.Add)
	r.HandleFunc("/sub", api.Subtraction)
	r.HandleFunc("/mul", api.Multiplication)
	r.HandleFunc("/div", api.Division)
	fmt.Printf("The server is started!!!\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
