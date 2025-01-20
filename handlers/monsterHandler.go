package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SergioPopovs176/dnd-library/app"
	"github.com/SergioPopovs176/dnd-library/storage"
)

type MonsterHandler struct {
	logger  *log.Logger
	storage storage.Storage
}

func NewMonsterHandler(a *app.Application) (*MonsterHandler, error) {
	return &MonsterHandler{
		logger:  a.Logger,
		storage: a.Storage,
	}, nil
}

func (h MonsterHandler) HandleGetMonstersList(w http.ResponseWriter, r *http.Request) {
	monsters, err := h.storage.GetMonsterList()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(monsters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (h MonsterHandler) HandleGetMonster(w http.ResponseWriter, r *http.Request) {
	monsterId := r.PathValue("id")
	fmt.Println(monsterId)

	id, err := strconv.Atoi(monsterId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	monster, err := h.storage.GetMonsterById(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(monster)

	resp, err := json.Marshal(monster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (h MonsterHandler) HandleAddMonster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add new monster\n"))
}

func (h MonsterHandler) HandleDeleteMonster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete monster by id\n"))
}

func (h MonsterHandler) HandleUpdateMonster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update monster by id\n"))
}
