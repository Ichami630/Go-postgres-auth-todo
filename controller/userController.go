package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ichami630/Go-JWT-Auth/config"
	"github.com/ichami630/Go-JWT-Auth/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Parse request body
	var body struct {
		Email           string `form:"email" json:"email" binding:"required"`
		Password        string `form:"password" json:"password" binding:"required"`
		ConfirmPassword string `form:"cpassword" json:"cpassword" binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input fields"})
		return
	}

	// Check if passwords match
	if body.Password != body.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user
	user := model.User{Email: body.Email, Password: string(passwordHash)}
	if result := config.Conn.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user. Email might already exist."})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{"success": "User registered successfully!"})
}

func Login(c *gin.Context) {
	// 	//get the emain and password from request
	var body struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fail to read body",
		})

		return
	}

	// 	//lookup requested user
	var user model.User
	config.Conn.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user or password",
		})
		return
	}

	// 	//compare the save and send password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user or password",
		})
		return
	}

	// 	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// // 	//sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fail to generate token",
		})
		return
	}

	//send the token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"success": tokenString,
	})
}
