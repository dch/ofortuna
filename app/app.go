package app

import (
	"fmt"
	"net/http"
	"ofortuna"
	"time"
)

func ListenAndServe(address string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, ofortuna.GetRandomFortune())
	})
	s := &http.Server{
		Addr: address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,			    Handler: mux,
	}
	return s.ListenAndServe()
}
