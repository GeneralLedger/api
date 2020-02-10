package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigratedDatabase(databaseURI string) (*sql.DB, error) {
	// Get a connection to the database
	db, err := sql.Open("postgres", databaseURI)
	if err != nil {
		return nil, err
	}

	// Run our migrations against the database connection
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	dbMigrator, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return nil, err
	}
	err = dbMigrator.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			// Eat the error & log here, we'll consider this a
			// non-error in this context.
			fmt.Println("Migrations up to date.")
		} else {
			return nil, err
		}
	}

	// OK
	return db, nil
}