package users

import (
	"errors"
	"events-app/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	// parsing user's POST JSON request
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// verifying the password entered by the user
	err = user.ValidateUser()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password!"})
		return
	}
	// generate JWT for the verified user
	token, err := GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token!"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Login successful!", "token": token})
}

// ValidateUser checks user's credentials
func (u *User) ValidateUser() error {
	// verify if user exists in the users table
	queryUserExists := `SELECT id, password FROM users WHERE email = $1;`
	row := db.DB.QueryRow(queryUserExists, u.Email)

	var dbPswd string
	err := row.Scan(&u.ID, &dbPswd)
	if err != nil {
		return errors.New("User does not exist")
	}

	pswdIsValid := CheckPswd(u.Password, dbPswd)
	if !pswdIsValid {
		return errors.New("incorrect password")
	}

	return nil
}
