package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var groceries = []Grocery{
	{Name: "Almod Milk", Quantity: 2},
	{Name: "Apple", Quantity: 6},
}

type Grocery struct {
	Name     string `json: "name"`
	Quantity int    `json: "quantity"`
}

func main() {
	fmt.Println("Hele")
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/allgroceries", AllGroceries)
	r.HandleFunc("/groceries/{name}", SingleGrocery)
	//r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")
	r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")
	r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", r))
}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)

			var updateGrocery Grocery

			json.NewDecoder(r.Body).Decode(&updateGrocery)
			groceries = append(groceries, updateGrocery)
			fmt.Println("Endpoint hit: UpdateGroceries")
			json.NewEncoder(w).Encode(updateGrocery)
			return
		}
	}
}

func SingleGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for _, grocery := range groceries {
		if grocery.Name == name {
			json.NewEncoder(w).Encode(grocery)
		}
	}
}

func AllGroceries(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllGroceries")
	json.NewEncoder(writer).Encode(groceries)
}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(
				groceries[:index],
				groceries[index+1:]...,
			)
		}
	}
}
