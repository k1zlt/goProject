package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Add(w http.ResponseWriter, r *http.Request) {
	paramA := r.URL.Query().Get("a")
	paramB := r.URL.Query().Get("b")
	a, err := strconv.Atoi(paramA)
	b, err2 := strconv.Atoi(paramB)
	if err != nil && err2 != nil {
		return
	}
	json.NewEncoder(w).Encode(a + b)
}
func Subtraction(w http.ResponseWriter, r *http.Request) {
	paramA := r.URL.Query().Get("a")
	paramB := r.URL.Query().Get("b")
	a, err := strconv.Atoi(paramA)
	b, err2 := strconv.Atoi(paramB)
	if err != nil && err2 != nil {
		return
	}
	json.NewEncoder(w).Encode(a - b)
}
func Multiplication(w http.ResponseWriter, r *http.Request) {
	paramA := r.URL.Query().Get("a")
	paramB := r.URL.Query().Get("b")
	a, err := strconv.Atoi(paramA)
	b, err2 := strconv.Atoi(paramB)
	if err != nil && err2 != nil {
		return
	}
	json.NewEncoder(w).Encode(a * b)
}
func Division(w http.ResponseWriter, r *http.Request) {
	paramA := r.URL.Query().Get("a")
	paramB := r.URL.Query().Get("b")
	a, err := strconv.Atoi(paramA)
	b, err2 := strconv.Atoi(paramB)
	if err != nil && err2 != nil {
		return
	}
	if b == 0 {
		json.NewEncoder(w).Encode("Cannot divide by zero")
		return
	}
	json.NewEncoder(w).Encode(a / b)
}
