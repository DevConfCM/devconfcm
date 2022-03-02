package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"devconfcm/pkg/config"
)

var db *gorm.DB

func init() {
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open(config.DefaultConfigs.TestDB), &gorm.Config{})
		if err != nil {
			panic("failed to connect to the database")
		}
	}
	db.AutoMigrate(&User{})
}
