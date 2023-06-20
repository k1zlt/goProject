package handler

import (
	"first/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/lessons/{id}", h.getLessonByID).Methods(http.MethodGet)
	r.HandleFunc("/lessons/{id}/video", h.getVideoByID).Methods(http.MethodGet)

	return r
}
