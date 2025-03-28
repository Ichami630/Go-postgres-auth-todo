package main

import (
	"github.com/ichami630/Go-JWT-Auth/config"
	"github.com/ichami630/Go-JWT-Auth/routes"
)

func main() {
	//load env variables
	config.LoadEnvVariables()

	//db connection
	config.Connection()

	//start the server
	routes.Router()
}
