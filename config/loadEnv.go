package config

import (
	"log"

	"github.com/joho/godotenv"
)

// load the env variables
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file ", err)
	}
}
