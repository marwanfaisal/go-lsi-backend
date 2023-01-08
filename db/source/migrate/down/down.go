package main

import (
	"log"

	"go-lsi-backend/db"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	defer func() {
		if err := db.SqlSourceDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	driver, err := postgres.WithInstance(db.SqlSourceDb, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrateDb, err := migrate.NewWithDatabaseInstance("file://db/source/migrate/sql", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err = migrateDb.Steps(-1); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}
