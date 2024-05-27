package users

import (
	"errors"
	"events-app/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int64
	Login     string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	FirstName string
	LastName  string
}

func SignUp(ctx *gin.Context) {
	// Parsing users' JSON POST request
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Unique()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"!message": "User login or email not unique!", "error": err.Error()})
		return
	}

	// create new user in the users DB
	err = user.New()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"!message": "Could not create user!", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User successfully created!"})

}

// New method creates new user in the users table
func (u *User) New() error {
	query := "INSERT INTO users (login, email, password) VALUES ($1, $2, $3)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatalf("Error preparing the query: v%\n", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Login, u.Email, u.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	// converting user's plain text password into hashed one

	// assign user id generated by the DB
	u.ID, _ = result.LastInsertId()

	return nil
}

// Unique method verifies if login and/or email exists in the users table
func (u *User) Unique() error {
	// check if users table contains the new user's login value
	queryLogin := `
	SELECT login
	FROM users
	WHERE login = $1;`

	rowsLogin, err := db.DB.Query(queryLogin, u.Login)
	if err != nil {
		return err
	}

	if rowsLogin.Next() {
		return errors.New("login not available")
	}

	// check if users table contains the new user's email value
	queryEmail := `
	SELECT email
	FROM users
	WHERE email = $1;`

	rowsEmail, err := db.DB.Query(queryEmail, u.Email)
	if err != nil {
		return err
	}

	if rowsEmail.Next() {
		return errors.New("user with this email already exists")
	}

	return nil
}
