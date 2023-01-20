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
	http.ListenAndServe(":3000", s.server)
}

var applicationSet = wire.NewSet(
	NewApplication,
	NewServer,
	NewRouteGroup,
	NewDatabase,
)

type Application struct {
	server *chi.Mux
	routes *chi.Router
	db     *gorm.DB
}

func NewApplication(server *chi.Mux, routes *chi.Router, db *gorm.DB) *Application {
	return &Application{
		server: server,
		routes: routes,
		db:     db,
	}
}

func NewServer() *chi.Mux {

	s := chi.NewRouter()

	return s
}

func NewRouteGroup(server *chi.Mux) *chi.Router {

	rg := server.Group(func(r chi.Router) {
		server.Get("/", GetMessage)
	})

	return &rg
}

func NewDatabase() *gorm.DB {
	return &gorm.DB{}
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	// a.db.First() // Query db with injected Database dependency
	w.Write([]byte("Number of the Day: 1"))
}
