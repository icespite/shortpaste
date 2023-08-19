package api

import (
	"github.com/timeforaninja/shortpaste/internal/types"
)

// App struct containing the bind address.
type app struct {
	bind            string
	link307Redirect bool
	fileDB          types.FileDB
	dataDB          types.DataDB
}

func (app *app) ShouldLink307Redirect() bool {
	return app.link307Redirect
}

func (app *app) GetFileDB() types.FileDB {
	return app.fileDB
}

func (app *app) GetDataDB() types.DataDB {
	return app.dataDB
}
