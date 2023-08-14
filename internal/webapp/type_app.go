package api

import (
	"github.com/timeforaninja/shortpaste/internal/data_db"
	"github.com/timeforaninja/shortpaste/internal/file_db"
)

// App struct containing the bind address.
type App struct {
	bind            string
	link307Redirect bool
	fileDB          *file_db.FileDB
	dataDB          *data_db.DataDB
}

func (app *App) ShouldLink307Redirect() bool {
	return app.link307Redirect
}

func (app *App) GetFileDB() *file_db.FileDB {
	return app.fileDB
}

func (app *App) GetDataDB() *data_db.DataDB {
	return app.dataDB
}
