package controller

import (
	"godog-demo/handler"
	"net/http"
)

func GetVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		handler.Fail(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := struct {
		Version string `json:"version"`
	}{Version: "0.0.1"}

	handler.Ok(w, data)
}
