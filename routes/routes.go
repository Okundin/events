package routes

import (
	"events-app/events"
	"events-app/middleware"
	"events-app/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// POST: HTTP REQUESTS
	server.POST("/signup", users.SignUp)
	server.POST("/login", users.Login)

	// GET: HTTP REQUESTS

	// PUT: HTTP REQUESTS

	// DELETE: HTTP REQUESTS

	// MIDDLEWARE
	auth := server.Group("/")
	auth.Use(middleware.Authenticate) // will be applied to all routes in the group
	// POST: HTTP REQUESTS
	auth.POST("/event-new", events.New)

	// GET: HTTP REQUESTS
	auth.GET("/event/all", events.GetAll)
	auth.GET("/event/:id", events.GetEvent)

	// PUT: HTTP REQUESTS
	auth.PUT("/user-update-name/:id", users.Update)
	auth.PUT("/event-update/:id", events.Update)

	// DELETE: HTTP REQUESTS
}
