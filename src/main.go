package main

import (
	"godog-demo/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/version", controller.GetVersion)
	http.ListenAndServe(":8080", nil)
}
