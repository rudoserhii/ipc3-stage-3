package db

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/obiMadu/ipc3-stage-3/internal/interfaces"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// new db
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to init db: %v", err)
	}

	if gin.Mode() == gin.ReleaseMode {
		db.Logger.LogMode(0)
	}

	DB = db

	rawDB := RawDB()

	rawDB.SetMaxIdleConns(20)
	rawDB.SetMaxOpenConns(100)

	// migrate models
	err = migrate()
	if err != nil {
		log.Panicf("Unable to migrate models %s\n", err.Error())
	}
	log.Println("Successfully Migrated Models.")

	return db
}

func migrate() error {
	err := DB.AutoMigrate(&interfaces.Product{}, &interfaces.Order{})
	if err != nil {
		return err
	}

	return nil
}

func RawDB() *sql.DB {
	rawDB, err := DB.DB()
	if err != nil {
		log.Panicf("Unable to get raw sql.DB %s\n", err.Error())
	}

	return rawDB
}
