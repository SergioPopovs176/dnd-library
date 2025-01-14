package app

import (
	"flag"
	"log"
	"os"
)

const version = "v0.0.1"

type config struct {
	Port    int
	Env     string
	Version string
}

type Application struct {
	Config config
	Logger *log.Logger
}

func New() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var cfg config

	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	cfg.Version = version

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	return app, nil
}
