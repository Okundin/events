package middleware

import (
	"events-app/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authenticate verifies the token provided by the user in the HTTP request
func Authenticate(ctx *gin.Context) {
	// here we parse the token provided in the request
	token := ctx.Request.Header.Get("Authorization")

	//case token is empty
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized!"})
		return
	}
	// case token is not valid; we also get a UserID value from the token to identify the user sending the HTTP request
	userID, err := users.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized!", "reason": err.Error(), "what to do?": "Please login again"})
		return
	}

	ctx.Set("userID", userID) //allows to use userID value everywhere where *gin.Context is present
	ctx.Next()

}
