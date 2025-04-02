package routes

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ichami630/Go-JWT-Auth/controller"
	"github.com/ichami630/Go-JWT-Auth/middleware"
	csrf "github.com/utrack/gin-csrf"
)

func Router() {
	server := gin.Default() //create a gin router with default middleware(logger and recovery)

	//enable csrf middleware
	store := cookie.NewStore([]byte(os.Getenv("CSRF_SECRET")))
	server.Use(sessions.Sessions("csrfsession", store))
	server.Use(csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
		ErrorFunc: func(c *gin.Context) {
			c.JSON(400, gin.H{"error": "CSRF Mismatch"})
			c.Abort()
		},
	}))

	server.Static("./assets", "./views/assets") //serve static files like css,js,images

	server.LoadHTMLGlob("views/*.html") //load html files from the views folder with support for nested folders

	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"csrfToken": csrf.GetToken(c)})
	})
	server.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{"csrfToken": csrf.GetToken(c)})
	})
	server.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", gin.H{"csrfToken": csrf.GetToken(c)})
	})

	server.GET("/admin", middleware.Auth, func(c *gin.Context) {
		c.HTML(200, "dashboard.html", gin.H{"csrfToken": csrf.GetToken(c)})
	})

	server.POST("/signup", controller.SignUp)

	server.POST("/login", controller.Login)

	// group admin route
	// admin := server.Group("/admin")
	// admin.Use(middleware.Auth)
	// {
	// 	admin.GET("/", func(c *gin.Context) {
	// 		c.HTML(200, "/dashboard.html", nil)
	// 	})
	// }

	server.Run(":8000") //listening on port 8080
}
