package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ichami630/Go-JWT-Auth/controller"
	"github.com/ichami630/Go-JWT-Auth/middleware"
)

func Router() {
	server := gin.Default() //create a gin router with default middleware(logger and recovery)

	server.Static("./assets", "./views/assets") //serve static files like css,js,images

	server.LoadHTMLGlob("views/*.html") //load html files from the views folder

	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	server.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	server.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})

	server.GET("/admin", middleware.Auth, func(c *gin.Context) {
		c.HTML(200, "dashboard.html", nil)
	})

	server.POST("/signup", controller.SignUp)

	server.POST("/login", controller.Login)

	//group admin route
	// admin := server.Group("/admin")
	// {
	// 	admin.GET("/", func(c *gin.Context) {
	// 		c.HTML(200, "dashboard.html", nil)
	// 	})
	// }

	server.Run(":8000") //listening on port 8080
}
