package types

import (
	"github.com/timeforaninja/shortpaste/internal/data_db"
	"github.com/timeforaninja/shortpaste/internal/file_db"
)

type AppInf interface {
	ShouldLink307Redirect() bool
	GetFileDB() *file_db.FileDB
	GetDataDB() *data_db.DataDB
}
