package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ichami630/Go-JWT-Auth/config"
	"github.com/ichami630/Go-JWT-Auth/model"
)

func Auth(c *gin.Context) {
	// Get the cookie from the request
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.Redirect(http.StatusFound, "/login") // Redirect to login page
		c.Abort()
		return
	}

	// Decode and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		c.Redirect(http.StatusFound, "/login") // Redirect to login page
		c.Abort()
		return
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	// Check expiration
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	// Find user
	var user model.User
	config.Conn.First(&user, claims["sub"])
	if user.ID == 0 {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	// Attach user to context and continue
	c.Set("user", user)
	c.Next()
}
