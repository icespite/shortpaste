package types

import "gorm.io/gorm"

type DataDB interface {
	AutoMigrate()
	GetUserPass() (string, string)
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
}
