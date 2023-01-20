package main

import "github.com/google/wire"

func InitializeApplication() *Application {
	wire.Build(
		applicationSet,
	)

	return nil
}
