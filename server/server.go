package server

import (
	"log"
	"net/http"
)

type Server struct {
	version int
}

func NewServer() (*Server, error) {
	return &Server{
		version: 1,
	}, nil
}

func (s Server) Start() {
	http.HandleFunc("/ping", handlePing)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
