package api

import (
	"github.com/timeforaninja/shortpaste/internal/types"
)

// App struct containing the bind address.
type App struct {
	bind            string
	link307Redirect bool
	fileDB          types.FileDB
	dataDB          types.DataDB
}

func (app *App) ShouldLink307Redirect() bool {
	return app.link307Redirect
}

func (app *App) GetFileDB() types.FileDB {
	return app.fileDB
}

func (app *App) GetDataDB() types.DataDB {
	return app.dataDB
}
