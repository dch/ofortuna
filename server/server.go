package server

import (
	"net/http"
)

func eServer(binding string, handler *func) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(server, nil)
}
