package main

import (
	"log"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/teclegacy/golang-ecom/config"
	db_ "github.com/teclegacy/golang-ecom/db"
)

func main() {

	db, err := db_.NewMySqlStorage(mysqlCfg.Config{
		User:   config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr:   config.Envs.DBAddress,
		DBName: config.Envs.DBName,

		AllowNativePasswords: true,
		ParseTime:            true,
		Net:                  "tcp",
	})
	if err != nil {
		log.Fatalf("Failed to create MySQL storage: %v", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create MySQL driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	cmd := os.Args[(len(os.Args) - 1)]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to rollback migrations: %v", err)
		}
	}
}
