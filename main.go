package main

import (
	"net/http"

	"go.uber.org/fx"

	"joshuapare.com/fxdb/internal/engine"
	"joshuapare.com/fxdb/internal/server"
	"joshuapare.com/fxdb/internal/storage"
)

func main() {
	fx.New(
		fx.Provide(
			server.NewHTTPServer,
			server.NewServeMux,
			storage.NewFSStorageEngine,
			engine.NewIndexEngine,
			engine.NewCollectionEngine,
			server.NewQueryHandler,
		),
		fx.Invoke(
			storage.SetupFS,
			func(*http.Server) {},
		),
	).Run()
}
