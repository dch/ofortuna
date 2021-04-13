package app

import (
	"fmt"
	"net/http"
	"ofortuna"
)

func ListenAndServe(address string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprintf(w, ofortuna.GetRandomFortune())
	})

	return http.ListenAndServe(address, nil)
}
