package handler

import (
	"encoding/json"
	"first/helper"
	"first/internal/domain"
	"first/logs"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, 0)
		helper.ErrorResponse(w, err.Error(), http.StatusBadRequest, 1)
		return
	}
	userID, err := h.services.CreateUser(user)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, 0)
		helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
		return
	}

	logs.Log(logs.Info, "", r.Method, r.URL.Path, userID)
	helper.Response(w, "SingUp successful")

}

type signInStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInStruct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, 0)
		helper.ErrorResponse(w, err.Error(), http.StatusBadRequest, 1)
		return
	}
	token, userID, err := h.services.GenerateToken(input.Username, input.Password)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, userID)
		helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
		return
	}

	response := map[string]string{
		"token": token,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logs.Log(logs.Error, err.Error(), r.Method, r.URL.Path, 0)
		helper.ErrorResponse(w, err.Error(), http.StatusInternalServerError, 1)
		return
	}

	logs.Log(logs.Info, "", r.Method, r.URL.Path, userID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
