package server

import (
	"fmt"
	"net/http"
	"time"
)

type pingHandler struct{}

func (pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/ping", pingHandler{})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}
