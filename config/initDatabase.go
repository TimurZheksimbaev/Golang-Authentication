package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectToDB(config *AppConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", 
		config.DatabaseHost, 
		config.DatabaseUsername, 
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return db, err
}