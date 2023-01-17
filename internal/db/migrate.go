package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Migrating database...")
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	fmt.Println("Database migrated successfully")

	return nil
}
