package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type LessonService struct {
}

func (h *Handler) getLessonByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lessonID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	lesson, err := h.services.GetLessonByID(lessonID)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	response, err := json.Marshal(lesson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) getVideoByID(w http.ResponseWriter, r *http.Request) {

}
