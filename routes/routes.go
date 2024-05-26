package routes

import (
	"events-app/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", users.SignUp)
}
