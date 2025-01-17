package postgres

import (
	"log"

	"github.com/SergioPopovs176/dnd-library/storage"
)

type ServerHandler struct {
	logger  *log.Logger
	storage storage.Storage
}
