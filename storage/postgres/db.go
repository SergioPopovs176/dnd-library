package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SergioPopovs176/dnd-library/storage"
	dnd5e "github.com/SergioPopovs176/dnd5-client/dnd-5e"
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

func (db *Db) Sync(c *dnd5e.Client) error {
	// TODO check is table 'monsters' exist and is empty
	tableName := "monsters"
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_schema = 'public'
			  AND table_name = $1
		);
	`

	var exists bool
	err := db.connection.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		//TODO write to logger
		fmt.Printf("Таблица %s существует.\n", tableName)
		return nil
	} else {
		//TODO write to logger
		fmt.Printf("Таблица %s не существует.\n", tableName)
	}

	//TODO create table
	createTableQuery := fmt.Sprintf(`
		CREATE TABLE %s (
			id SERIAL NOT NULL UNIQUE,
			index varchar(255) NOT NULL UNIQUE,
			name varchar(255) NOT NULL,
			size varchar(255) NOT NULL,
			type varchar(255) NOT NULL,
			alignment varchar(255) NOT NULL,
			added_at timestamp NOT NULL DEFAULT now(),
			updated_at timestamp NOT NULL DEFAULT now()
		);
	`, tableName)

	r, err := db.connection.Exec(createTableQuery)
	fmt.Println(r)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы %s: %v", tableName, err)
	}

	fmt.Printf("Таблица %s успешно создана\n", tableName)

	monsters, err := c.GetMonsters()
	if err != nil {
		return err
	}
	fmt.Println(monsters)

	for _, m := range monsters {
		mf, err := c.GetMonster(m.Index)
		if err != nil {
			return err
		}

		dbm := storage.MonsterFull{
			Index:     mf.Index,
			Name:      mf.Name,
			Size:      mf.Size,
			Type:      mf.Type,
			Alignment: mf.Alignment,
		}

		_, err = db.connection.Exec("INSERT INTO monsters (index, name, size, type, alignment) values ($1, $2, $3, $4, $5)",
			dbm.Index, dbm.Name, dbm.Size, dbm.Type, dbm.Alignment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Db) GetMonsterList() ([]storage.Monster, error) {
	rows, err := db.connection.Query("select id, index from monsters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	monsters := make([]storage.Monster, 0)
	for rows.Next() {
		m := storage.Monster{}
		err := rows.Scan(&m.ID, &m.Index)
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

func (db *Db) GetMonsterById(monsterId int) (storage.MonsterFull, error) {
	m := storage.MonsterFull{}
	err := db.connection.QueryRow("select id, index, name, size, type, alignment from monsters where id = $1", monsterId).
		Scan(&m.ID, &m.Index, &m.Name, &m.Size, &m.Type, &m.Alignment)

	return m, err
}

func (db *Db) AddMonster(monster storage.MonsterFull) (storage.MonsterFull, error) {
	query := `
		INSERT INTO monsters (index, name, size, type, alignment)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	var id int
	err := db.connection.QueryRow(query, monster.Index, monster.Name, monster.Size, monster.Type, monster.Alignment).Scan(&id)
	if err != nil {
		log.Fatalf("Ошибка при добавлении записи: %v", err)
	}

	return db.GetMonsterById(id)
}

func (db *Db) DeleteMonsterById(monsterId int) error {
	_, err := db.connection.Exec("delete from monsters where id = $1", monsterId)
	return err
}

func (db *Db) UpdateMonsterById(monsterId int, monster storage.MonsterFull) error {
	_, err := db.connection.Exec("UPDATE monsters SET name=$1, size=$2, type=$3, alignment=$4 where id = $5",
		monster.Name, monster.Size, monster.Type, monster.Alignment, monsterId)
	fmt.Println(err)
	return err
}
