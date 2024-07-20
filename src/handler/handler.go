package handler

import (
	"encoding/json"
	"net/http"
)

func Fail(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)

	data := struct {
		Error string `json:"error"`
	}{Error: msg}
	resp, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func Ok(w http.ResponseWriter, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		Fail(w, "Oops something evil has happened", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}