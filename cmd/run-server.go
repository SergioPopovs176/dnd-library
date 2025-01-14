package main

import (
	"log"

	"github.com/SergioPopovs176/dnd-library/app"
	"github.com/SergioPopovs176/dnd-library/server"
)

// go run ./cmd/run-server.go
func main() {
	app, _ := app.New()

	app.Logger.Println("Start server on port", app.Config.Port, "-- Version", app.Config.Version)

	server, err := server.NewServer(app)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}

//TODO
//1) подготовить роуты
//2) использовать базу
//3) функция подгрузки данных из внешнего сервера
//4) docker
