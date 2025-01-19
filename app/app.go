package app

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SergioPopovs176/dnd-library/storage"
	"github.com/SergioPopovs176/dnd-library/storage/postgres"
	dnd5e "github.com/SergioPopovs176/dnd5-client/dnd-5e"
	"github.com/joho/godotenv"
)

type config struct {
	Port    int
	Env     string
	Version string
}

type Application struct {
	Config    config
	Logger    *log.Logger
	Storage   storage.Storage
	DndClient *dnd5e.Client
}

func New() (*Application, error) {
	// Load envs from file .env
	err := godotenv.Load()
	if err != nil {
		return &Application{}, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var cfg config

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return &Application{}, err
	}

	flag.IntVar(&cfg.Port, "port", port, "API server port")
	flag.StringVar(&cfg.Env, "env", os.Getenv("APP_ENV"), "Environment (dev|stage|prod)")
	flag.Parse()

	cfg.Version = os.Getenv("APP_VERSION")

	storage, err := postgres.NewStorage()
	if err != nil {
		return &Application{}, err
	}

	client, err := dnd5e.NewClient(time.Second * 15)
	if err != nil {
		log.Fatal(err)
	}

	app := &Application{
		Config:    cfg,
		Logger:    logger,
		Storage:   storage,
		DndClient: client,
	}

	return app, nil
}

func (app *Application) Stop() {
	app.Logger.Println("Start application closing. Need close conection to DB")
	app.Storage.Close()
}
