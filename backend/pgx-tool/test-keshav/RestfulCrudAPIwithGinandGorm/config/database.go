package config

import (
	"CrudAPIWithGin/helper"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "khushbu"
	dbName   = "ginDemoDB"
	//SSLMode  = "disable"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s  dbname=%s",
		host, port, password, user, dbName,
	)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
