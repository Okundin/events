package routes

import (
	"events-app/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", users.SignUp)
	server.POST("/login", users.Login)
	server.PUT("/user-update-name/:id", users.Update)
}
