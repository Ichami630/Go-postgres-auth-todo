package config

import "github.com/ichami630/Go-JWT-Auth/model"

func DbMigrations() {
	//migrate the user's table
	Conn.AutoMigrate(&model.User{})
}
