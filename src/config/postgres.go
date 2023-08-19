package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgre struct {
	//DB Configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string
}

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("database")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
