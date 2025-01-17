package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SergioPopovs176/dnd-library/app"
	"github.com/SergioPopovs176/dnd-library/handlers"
)

// curl -v -X GET -H 'x-id:4' localhost:8080/users

type Server struct {
	version        int
	port           int
	app            *app.Application
	monsterHandler *handlers.MonsterHandler
	serverHandler  *handlers.ServerHandler
}

func NewServer(app *app.Application) (*Server, error) {
	mh, err := handlers.NewMonsterHandler(app)
	if err != nil {
		return nil, err
	}

	sh, err := handlers.NewServerHandler(app)
	if err != nil {
		return nil, err
	}

	return &Server{
		version:        1,
		port:           app.Config.Port,
		app:            app,
		monsterHandler: mh,
		serverHandler:  sh,
	}, nil
}

func (s *Server) Start() {
	mux := getRouter(s)
	wrappedMux := loggerMiddleware(mux, s.app.Logger)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      wrappedMux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getRouter(s *Server) *http.ServeMux {
	mux := http.NewServeMux()

	// curl -v -X GET localhost:8080/v1/ping
	mux.HandleFunc("GET /v1/ping", s.serverHandler.HandlePing)
	// curl -v -X GET localhost:8080/v1/monsters
	mux.HandleFunc("GET /v1/monsters", s.monsterHandler.HandleGetMonstersList)
	// curl -v -X GET localhost:8080/v1/monsters/{id}
	mux.HandleFunc("GET /v1/monsters/{id}", s.monsterHandler.HandleGetMonster)
	// curl -v -X POST localhost:8080/v1/monsters/{id}
	mux.HandleFunc("POST /v1/monsters", s.monsterHandler.HandleAddMonster)
	// curl -v -X DELETE localhost:8080/v1/monsters/{id}
	mux.HandleFunc("DELETE /v1/monsters/{id}", s.monsterHandler.HandleDeleteMonster)
	// curl -v -X PUT localhost:8080/v1/monsters/{id}
	mux.HandleFunc("PUT /v1/monsters/{id}", s.monsterHandler.HandleUpdateMonster)

	// mux.HandleFunc("GET /v0/game/status", app.Game.GetStatusHandler)
	// mux.HandleFunc("POST /v0/game/add", app.Game.AddPlayerHandler)

	return mux
}

func loggerMiddleware(next http.Handler, l *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// idFromCtx := r.Context().Value("id")
		// userID, ok := idFromCtx.(string)
		// if !ok {
		// 	log.Printf("[%s] %s - error: userID is invalid\n", r.Method, r.URL)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		userID := "88"
		l.Printf("[%s] %s by userID %s\n", r.Method, r.URL, userID)
		next.ServeHTTP(w, r)
	})
}
