package data_db

import (
	"github.com/timeforaninja/shortpaste/internal/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
)

// DataDB struct containing the storage path and the db connector.
type DataDB struct {
	storagePath string
	db          *gorm.DB
	user        string
	password    string
}

func NewDataDB(storagePath, user, password string) (*DataDB, error) {
	dbUri := path.Join(storagePath, "db", "mapping.db")
	db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DataDB{
		storagePath: storagePath,
		db:          db,
		user:        user,
		password:    password,
	}, nil
}

func (db *DataDB) AutoMigrate() {
	db.db.AutoMigrate(&types.Link{})
	db.db.AutoMigrate(&types.File{})
	db.db.AutoMigrate(&types.Text{})
}

func (db *DataDB) GetUserPass() (string, string) {
	return db.user, db.password
}

func (db *DataDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.db.Find(dest, conds)
}

func (db *DataDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.db.First(dest, conds)
}

func (db *DataDB) Save(value interface{}) *gorm.DB {
	return db.db.Save(value)
}

func (db *DataDB) Create(value interface{}) *gorm.DB {
	return db.db.Create(value)
}
