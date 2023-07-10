package handler

import (
	"encoding/json"
	"first/helper"
	"first/logs"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) getLessonByID(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		logs.Log(logs.Error, "invalid user_id", r.Method, r.URL.Path, 0)
		helper.ErrorResponse(w, "invalid user id", http.StatusInternalServerError, 1)
		return
	}

	params := mux.Vars(r)
	lessonID, err := strconv.Atoi(params["id"])
	if err != nil {
		logs.Log(logs.Error, "invalid lesson_id", r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, "invalid lesson_id", http.StatusBadRequest, 1)
		return
	}

	lesson, err := h.services.GetLessonByID(userID, lessonID, r.URL.Path)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
		return
	}

	hasPermissionToLesson, err := h.services.IsLessonAccessibleForUser(userID, lessonID)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
	}
	if !hasPermissionToLesson {
		logs.Log(logs.Error, "user has no permission to this lesson", r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, "user has no permission to this lesson", http.StatusUnauthorized, 1)
		return
	}

	response, err := json.Marshal(lesson)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
		return
	}

	logs.Log(logs.Info, "", r.Method, r.URL.Path, userID)
	helper.Response(w, string(response))
}
