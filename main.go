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
	// add route group dep
	db *gorm.DB
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

func NewRouteGroup(app *Application) *chi.Router {

	rg := app.server.Group(func(r chi.Router) {
		r.Get("/", app.GetMessage)
	})

	return &rg

}

func (a *Application) GetMessage(w http.ResponseWriter, r *http.Request) {
	// a.db.First() // Query db with injected Database dependency
	w.Write([]byte("Number of the Day: 1"))
}
