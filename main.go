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
	http.ListenAndServe(":3000", s.router)
}

type Application struct {
	router *chi.Mux
	db     *gorm.DB
}

type Route struct {
	url    string
	method string
	fn     http.HandlerFunc
}

var applicationSet = wire.NewSet(
	NewDatabase,
	NewMessageRoute,
	NewMux,
	NewApplication,
)

func NewApplication(router *chi.Mux, db *gorm.DB) *Application {
	return &Application{
		router: router,
		db:     db,
	}
}

func NewMux(routes *[]Route) *chi.Mux {

	s := chi.NewRouter()

	s.Group(func(r chi.Router) {
		for _, route := range *routes {
			switch route.method {
			case "GET":
				s.Get(route.url, route.fn)
			}
		}
	})

	return s
}

func NewMessageRoute(db *gorm.DB) *[]Route {

	getMessage := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Number of the Day: 1"))
	}

	getMessageRoute := Route{
		url:    "/",
		method: "GET",
		fn:     getMessage,
	}

	return &[]Route{
		getMessageRoute,
	}
}

// func (app *Application) NewMessageRoute(router *chi.Router) {}

/*
	func getMessage(w http.ResponseWriter, r *http.Request) {
		// a.db.First() // Query db with injected Database dependency
		w.Write([]byte("Number of the Day: 1"))
	}
*/
func NewDatabase() *gorm.DB {
	return &gorm.DB{}
}
