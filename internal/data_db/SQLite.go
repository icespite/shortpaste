package data_db

import (
	"github.com/timeforaninja/shortpaste/internal/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
)

// SQLiteDB struct containing the storage path and the db connector for sqlite.
type sqliteDB struct {
	storagePath string
	db          *gorm.DB
	user        string
	password    string
}

func NewSQLiteDataDB(storagePath, user, password string) (types.DataDB, error) {
	dbUri := path.Join(storagePath, "db", "mapping.db")
	db, err := gorm.Open(sqlite.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &sqliteDB{
		storagePath: storagePath,
		db:          db,
		user:        user,
		password:    password,
	}, nil
}

func (db *sqliteDB) AutoMigrate() {
	db.db.AutoMigrate(&types.Link{})
	db.db.AutoMigrate(&types.File{})
	db.db.AutoMigrate(&types.Text{})
}

func (db *sqliteDB) GetUserPass() (string, string) {
	return db.user, db.password
}

func (db *sqliteDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.db.Find(dest, conds)
}

func (db *sqliteDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.db.First(dest, conds)
}

func (db *sqliteDB) Save(value interface{}) *gorm.DB {
	return db.db.Save(value)
}

func (db *sqliteDB) Create(value interface{}) *gorm.DB {
	return db.db.Create(value)
}
