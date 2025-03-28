package routes

import "github.com/gin-gonic/gin"

func Router() {
	server := gin.Default() //create a gin router with default middleware(logger and recovery)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "welcome"})
	})

	server.Run(":8080") //listening on port 8080
}
