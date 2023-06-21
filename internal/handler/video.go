package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) getVideoByLessonID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lessonID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	video, err := h.services.GetVideoByLessonID(lessonID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", strconv.Itoa(len(video.Data)))

	_, err = w.Write(video.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
