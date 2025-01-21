package main

import (
	"fmt"
	"log"

	"github.com/SergioPopovs176/dnd-library/app"
	"github.com/SergioPopovs176/dnd-library/server"
)

// go run ./cmd/run-server.go
func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Stop()

	//TODO move ping to storage ??
	if err := app.Storage.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONECTED")

	//TODO sync DB data
	err = app.Storage.Sync(app.DndClient)
	if err != nil {
		log.Fatal(err)
	}

	app.Logger.Println("Start server on port", app.Config.Port, "-- Version", app.Config.Version)

	server, err := server.NewServer(app)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
