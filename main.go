package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
)

func main() {
	s := InitializeApplication()

	s.Start(s)
	//http.ListenAndServe(":3000", s.server)
}

var applicationSet = wire.NewSet(
	NewApplication,
	NewServer,
	NewDatabase,
)

type Application struct {
	server *chi.Mux
	db     *gorm.DB
}

func NewApplication(server *chi.Mux, db *gorm.DB) *Application {
	return &Application{
		server: server,
		db:     db,
	}
}

func NewServer() *chi.Mux {

	s := chi.NewRouter()

	return s
}

func NewDatabase() *gorm.DB {
	return &gorm.DB{}
}

func (a *Application) Start(s *Application) {

	a.server.Get("/", a.GetMessage)

	http.ListenAndServe(":3000", s.server)
}

func (a *Application) GetMessage(w http.ResponseWriter, r *http.Request) {
	// a.db.First() // Query db with injected Database dependency
	w.Write([]byte("Number of the Day: 1"))
}
