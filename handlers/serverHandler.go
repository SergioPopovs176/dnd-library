package handlers

import (
	"log"
	"net/http"

	"github.com/SergioPopovs176/dnd-library/app"
	"github.com/SergioPopovs176/dnd-library/storage"
)

type ServerHandler struct {
	logger  *log.Logger
	storage storage.Storage
}

func NewServerHandler(a *app.Application) (*ServerHandler, error) {
	return &ServerHandler{
		logger:  a.Logger,
		storage: a.Storage,
	}, nil
}

func (h *ServerHandler) HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
