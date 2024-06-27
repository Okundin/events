package users

import (
	"events-app/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	// parsing user PUT HTTP request
	userID, err := strconv.ParseInt(ctx.Param("id"), 10, 64) // 10: decimal, 64: int64 type
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get UserID of the authenticated user from the token to compare it with the userID who owns the account
	// to make sure that only account owner can update his or her data
	authUserID := ctx.GetInt64("userID")
	if authUserID != userID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized!"})
		return
	}

	// verify if user exists in the users table
	user, err := GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Could not find user!"})
		return
	}
	// parsing body of the user HTTP request
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.updateName()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update name", "reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Success": "First and last name updated!"})
}

func GetUserByID(userID int64) (UserUpdate, error) {
	// verify if user requesting name update exists in the users DB table
	queryUser := `SELECT id FROM users WHERE id = $1;`
	row := db.DB.QueryRow(queryUser, userID)
	var u UserUpdate
	err := row.Scan(&u.ID)
	if err != nil {
		return UserUpdate{}, err
	}

	return u, nil
}

func (u *UserUpdate) updateName() error {
	query := `UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3;`
	_, err := db.DB.Exec(query, u.FirstName, u.LastName, u.ID)
	if err != nil {
		return err
	}

	return nil
}
