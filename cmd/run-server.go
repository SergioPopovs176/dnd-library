package main

import (
	"fmt"
	"log"

	"github.com/SergioPopovs176/dnd-library/server"
)

func main() {
	fmt.Println("Start server v0.0.1")

	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
