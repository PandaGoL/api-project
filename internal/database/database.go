package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	fmt.Println("Connection to DB")
	dbUsername := "postgress"
	dbPassword := "12345"
	dbHost := "localhost"
	dbTable := "postgres"
	dbPort := "5432"

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		return db, err
	}

	return db, nil
}
