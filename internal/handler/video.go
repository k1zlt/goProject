package handler

import (
	"first/helper"
	"first/logs"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) getVideoByLessonID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		logs.Log(logs.Error, "invalid user_id", r.Method, r.URL.Path, 0)
		helper.ErrorResponse(w, "invalid user_id", http.StatusInternalServerError, 1)
		return
	}

	params := mux.Vars(r)
	lessonID, err := strconv.Atoi(params["id"])
	if err != nil {
		logs.Log(logs.Error, "invalid lesson_id", r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, "invalid lesson_id", http.StatusBadRequest, 0)
		return
	}

	hasPermissionToLesson, err := h.services.IsLessonAccessibleForUser(userID, lessonID)
	if err != nil {
		logs.Log(logs.Error, "", r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
	}
	if !hasPermissionToLesson {
		logs.Log(logs.Error, "user has no permission to this lesson", r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, "user has no permission to this lesson", http.StatusUnauthorized, 1)
		return
	}

	video, err := h.services.GetVideoByLessonID(lessonID)
	if err != nil {
		logs.Log(logs.Error, "no lesson with indicated id", r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, "no lesson with indicated id", http.StatusInternalServerError, 1)
		return
	}

	logs.Log(logs.Info, "", r.Method, r.URL.Path, userID)
	helper.ResponseVideo(w, r, userID, video.Data)

}
