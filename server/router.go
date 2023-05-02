package server

import (
	controllers "notify-email/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Routes
	r.GET("/health", controllers.Health)
}
