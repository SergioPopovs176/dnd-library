package storage

type Monster struct {
	ID int
}

// Index     string `json:"index"`
// Name      string `json:"name"`
// Url       string `json:"url"`
// Size      string `json:"size"`
// Type      string `json:"type"`
// Subtype   string `json:"subtype"`
// Alignment string `json:"alignment"`

type Storage interface {
	Close()
	Ping() error
	Sync() error
	GetMonsterList() ([]Monster, error)
}
