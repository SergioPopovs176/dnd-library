package storage

import dnd5e "github.com/SergioPopovs176/dnd5-client/dnd-5e"

type Monster struct {
	ID    int
	Index string `json:"index"`
}

type MonsterFull struct {
	ID        int
	Index     string `json:"index"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	Type      string `json:"type"`
	Alignment string `json:"alignment"`
}

type Storage interface {
	Close()
	Ping() error
	Sync(*dnd5e.Client) error
	GetMonsterList() ([]Monster, error)
	GetMonsterById(monsterId int) (MonsterFull, error)
}
