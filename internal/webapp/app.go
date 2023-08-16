package api

import (
	"fmt"
	"github.com/timeforaninja/shortpaste/internal/data_db"
	"github.com/timeforaninja/shortpaste/internal/file_db"
	"github.com/timeforaninja/shortpaste/internal/routes"
	"github.com/timeforaninja/shortpaste/internal/utils"
	"os"
	"path"
)

// Run migrates the DB and starts the web server
func (app *App) Run() {
	app.GetDataDB().AutoMigrate()
	fmt.Println("Migration complete")
	routes.HandleRequests(app, app.bind)
}

// NewApp creates a new App instance with the provided bind address and storage path
func NewApp(bind, storagePath, username, password string, link307Redirect bool) App {
	storagePath = utils.EscapeHomePath(storagePath)
	err := os.MkdirAll(path.Join(storagePath, "db"), 0700)
	if err != nil {
		panic(fmt.Errorf("mkdir error %v", err))
	}

	ddb, err := data_db.NewSQLiteDataDB(storagePath, username, password)
	if err != nil {
		panic(fmt.Errorf("data db error %v", err))
	}

	fdb := file_db.NewLocalFileDB(storagePath)

	return App{
		bind:            bind,
		dataDB:          ddb,
		fileDB:          fdb,
		link307Redirect: link307Redirect,
	}
}
