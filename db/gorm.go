package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SourceDB *gorm.DB
var DestinationDB *gorm.DB

func Open() (gormSourceDb, gormDestinationDb *gorm.DB) {
	var err error
	if gormSourceDb, err = gorm.Open(postgres.New(postgres.Config{Conn: SqlSourceDb}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		log.Fatal(err)
		return
	}
	SourceDB = gormSourceDb

	if gormDestinationDb, err = gorm.Open(postgres.New(postgres.Config{Conn: SqlDestinationDb}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		log.Fatal(err)
		return
	}
	DestinationDB = gormDestinationDb
	return
}
