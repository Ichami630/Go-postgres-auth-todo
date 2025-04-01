package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ichami630/Go-JWT-Auth/config"
	"github.com/ichami630/Go-JWT-Auth/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//get the emain and password from request
	var body struct {
		email           string
		password        string
		confirmPassword string
	}

	if c.Bind(body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fail to read body",
		})

		return
	}
	//check password match
	if body.password != body.confirmPassword {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"msg": "passwords do not match",
		})

		return
	}

	//hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "fail to hash password",
		})
		return
	}

	//create the user
	user := model.User{Email: body.email, Password: string(passwordHash)}
	if result := config.Conn.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
	}

	//give the response
	c.HTML(http.StatusBadRequest, "signup.html", gin.H{
		"msg": "user created successfully",
	})
}
