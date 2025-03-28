package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB //global conn variable

func Connection() {
	var err error

	if Conn, err = gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{}); err != nil {
		log.Fatal("Database connection error:", err)
	}

}
