// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitializeApplication() *Application {
	db := NewDatabase()
	v := NewMessageRoute(db)
	mux := NewMux(v)
	application := NewApplication(mux, db)
	return application
}
