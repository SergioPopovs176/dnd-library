package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/SergioPopovs176/dnd-library/storage"
	_ "github.com/lib/pq"
)

type config struct {
	host     string
	port     string
	user     string
	dbName   string
	sslMode  string
	password string
}

type Db struct {
	version    string
	connection *sql.DB
}

func NewStorage() (*Db, error) {
	cfg := config{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		dbName:   os.Getenv("DB_NAME"),
		sslMode:  os.Getenv("DB_SSL_MODE"),
		password: os.Getenv("DB_PASSWORD"),
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.host, cfg.port, cfg.user, cfg.dbName, cfg.sslMode, cfg.password)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return &Db{}, err
	}

	return &Db{version: "v1", connection: db}, nil
}

func (db *Db) Close() {
	db.connection.Close()
}

func (db *Db) Ping() error {
	return db.connection.Ping()
}

func (db *Db) Sync() error {
	// TODO check is table 'monsters' exist and is empty
	// Проверка существования таблицы

	// tableName := "your_table_name"
	// query := `
	// 	SELECT EXISTS (
	// 		SELECT 1
	// 		FROM information_schema.tables
	// 		WHERE table_schema = 'public'
	// 		  AND table_name = $1
	// 	);
	// `

	// var exists bool
	// err = db.QueryRow(query, tableName).Scan(&exists)
	// if err != nil {
	// 	log.Fatalf("Ошибка выполнения запроса: %v", err)
	// }

	// if exists {
	// 	fmt.Printf("Таблица %s существует.\n", tableName)
	// } else {
	// 	fmt.Printf("Таблица %s не существует.\n", tableName)
	// }

	return nil
}

func (db *Db) GetMonsterList() ([]storage.Monster, error) {
	rows, err := db.connection.Query("select * from monsters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	monsters := make([]storage.Monster, 0)
	for rows.Next() {
		m := storage.Monster{}
		err := rows.Scan(&m.ID)
		if err != nil {
			return nil, err
		}

		monsters = append(monsters, m)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return monsters, nil
}
