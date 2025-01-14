package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SergioPopovs176/dnd-library/app"
)

// curl -v -X GET -H 'x-id:4' localhost:8080/users

type Server struct {
	version int
	port    int
}

func NewServer(app *app.Application) (*Server, error) {
	return &Server{
		version: 1,
		port:    app.Config.Port,
	}, nil
}

func (s Server) Start() {
	mux := getRouter()
	wrappedMux := loggerMiddleware(mux)
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

func getRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// curl -v -X GET localhost:8080/v1/ping
	mux.HandleFunc("GET /v1/ping", handlePing)
	// curl -v -X GET localhost:8080/v1/monsters
	mux.HandleFunc("GET /v1/monsters", handleGetMonstersList)
	// curl -v -X GET localhost:8080/v1/monsters/{id}
	mux.HandleFunc("GET /v1/monsters/{id}", handleGetMonster)
	// curl -v -X POST localhost:8080/v1/monsters/{id}
	mux.HandleFunc("POST /v1/monsters", handleAddMonster)
	// curl -v -X DELETE localhost:8080/v1/monsters/{id}
	mux.HandleFunc("DELETE /v1/monsters/{id}", handleDeleteMonster)
	// curl -v -X PUT localhost:8080/v1/monsters/{id}
	mux.HandleFunc("PUT /v1/monsters/{id}", handleUpdateMonster)

	// mux.HandleFunc("GET /v0/game/status", app.Game.GetStatusHandler)
	// mux.HandleFunc("POST /v0/game/add", app.Game.AddPlayerHandler)

	return mux
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// idFromCtx := r.Context().Value("id")
		// userID, ok := idFromCtx.(string)
		// if !ok {
		// 	log.Printf("[%s] %s - error: userID is invalid\n", r.Method, r.URL)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		userID := "88"
		log.Printf("[%s] %s by userID %s\n", r.Method, r.URL, userID)
		next.ServeHTTP(w, r)
	})
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

func handleGetMonstersList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list\n"))
}

func handleGetMonster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("monster full info by id\n"))
}

func handleAddMonster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Add new monster\n"))
}

func handleDeleteMonster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete monster by id\n"))
}

func handleUpdateMonster(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update monster by id\n"))
}
