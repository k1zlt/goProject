package helper

import (
	"encoding/json"
	"first/internal/domain"
	"first/logs"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, text string, errorCode, statusCode int) {
	resp := &domain.Response{
		Message:       "",
		ErrorResponse: text,
		ErrorCode:     statusCode,
	}
	w.WriteHeader(errorCode)
	data, err := json.Marshal(resp)
	if err != nil {
		logs.Log(logs.Fatal, err.Error(), "", "", 0)
	}
	w.Write(data)

}
func Response(w http.ResponseWriter, param string) {
	resp := &domain.Response{
		Message:       param,
		ErrorResponse: "",
		ErrorCode:     0,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		logs.Log(logs.Fatal, err.Error(), "", "", 0)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ResponseVideo(w http.ResponseWriter, r *http.Request, userID int, video []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(video)

}
