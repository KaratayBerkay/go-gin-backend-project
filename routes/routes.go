package routes

import (
	"github.com/KaratayBerkay/go-gin-backend-project/controllers"
	"github.com/KaratayBerkay/go-gin-backend-project/middleware"
	"github.com/gin-gonic/gin"
)


// SetupRouter configures the API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Create controller instances
	pingController := &controllers.PingController{}

	// Public routes
	r.GET("/ping", pingController.Ping)

	// Private routes - require authentication
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Add your protected routes here
		// Example: authorized.GET("/users", userController.GetUsers)
	}

	return r
}
