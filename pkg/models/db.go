package models

import (
	"devconfcm/pkg/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open(config.DefaultConfig.TestDB), &gorm.Config{})
		if err != nil {
			panic("failed to connect to the database")
		}
	}
	db.AutoMigrate(&User{})
}
