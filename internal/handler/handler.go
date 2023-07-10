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

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)
	auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)

	lessons := r.PathPrefix("/lessons").Subrouter()
	lessons.Use(h.authorizeUser, h.checkEndpointPermission)
	lessons.HandleFunc("/{id}", h.getLessonByID).Methods(http.MethodGet)
	lessons.HandleFunc("/{id}/video", h.getVideoByLessonID).Methods(http.MethodGet)

	return r
}
