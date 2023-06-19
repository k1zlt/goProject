package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

var connectStr = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func GetLesson(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatal(err)
	}

	query_id_ := r.URL.Query().Get("id")
	query_id, err := strconv.Atoi(query_id_)
	query := "SELECT * FROM lesson"
	rows, err := db.Query(query)
	for rows.Next() {
		var id int
		var title string
		var content string
		var cotegory_id int

		err := rows.Scan(&id, &title, &content, &cotegory_id)
		if err != nil {

		}
		if id == query_id {
			json.NewEncoder(w).Encode(fmt.Sprintf("%s\n%s", title, content))
			return
		}
	}
	json.NewEncoder(w).Encode("NOT FOUND 404")
}

//func Add(w http.ResponseWriter, r *http.Request) {
//	paramA := r.URL.Query().Get("a")
//	paramB := r.URL.Query().Get("b")
//	a, err := strconv.Atoi(paramA)
//	b, err2 := strconv.Atoi(paramB)
//	if err != nil && err2 != nil {
//		return
//	}
//	json.NewEncoder(w).Encode(a + b)
//}
//func Subtraction(w http.ResponseWriter, r *http.Request) {
//	paramA := r.URL.Query().Get("a")
//	paramB := r.URL.Query().Get("b")
//	a, err := strconv.Atoi(paramA)
//	b, err2 := strconv.Atoi(paramB)
//	if err != nil && err2 != nil {
//		return
//	}
//	json.NewEncoder(w).Encode(a - b)
//}
//func Multiplication(w http.ResponseWriter, r *http.Request) {
//	paramA := r.URL.Query().Get("a")
//	paramB := r.URL.Query().Get("b")
//	a, err := strconv.Atoi(paramA)
//	b, err2 := strconv.Atoi(paramB)
//	if err != nil && err2 != nil {
//		return
//	}
//	json.NewEncoder(w).Encode(a * b)
//}
//func Division(w http.ResponseWriter, r *http.Request) {
//	paramA := r.URL.Query().Get("a")
//	paramB := r.URL.Query().Get("b")
//	a, err := strconv.Atoi(paramA)
//	b, err2 := strconv.Atoi(paramB)
//	if err != nil && err2 != nil {
//		return
//	}
//	if b == 0 {
//		json.NewEncoder(w).Encode("Cannot divide by zero")
//		return
//	}
//	json.NewEncoder(w).Encode(a / b)
//}
