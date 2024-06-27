package users

import (
	"errors"
	"events-app/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	// Parsing users' JSON POST request
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// bring login and email strings entered by new user to lower case
	user.Login = strings.ToLower(user.Login)
	user.Email = strings.ToLower(user.Email)

	// check if login and email are unique before posting to the table
	err = user.Unique()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"!message": "User login or email not unique!", "error": err.Error()})
		return
	}

	// create new user in the users table assigning default role 'user'
	user.Role = "user"
	err = user.New()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"!message": "Could not create user!", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"1 Congrats!": "User successfully created!", "2 Login": user.Login, "3 Email": user.Email,
		"4 UserID": user.ID, "5 Created at": user.CreatedAt})
}

// New method creates new user in the users table
func (u *User) New() error {
	// converting user's plain text password into hashed one
	hashedPswd := HashPswd(u.Password)

	// saving new user's data to the users table
	query := "INSERT INTO users (login, email, password, created_at, role) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, $4) RETURNING id, created_at;"

	err := db.DB.QueryRow(query, u.Login, u.Email, hashedPswd, u.Role).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// Unique method verifies if login and/or email exists in the users table
func (u *User) Unique() error {
	// check if users table contains the new user's login value
	queryLogin := `SELECT login FROM users WHERE login = $1;`

	rowsLogin, err := db.DB.Query(queryLogin, u.Login)
	if err != nil {
		return err
	}
	// returns true if login already exists in the table, or false if not exists
	if rowsLogin.Next() {
		return errors.New("login not available")
	}

	// check if users table contains the new user's email value
	queryEmail := `SELECT email FROM users WHERE email = $1;`

	rowsEmail, err := db.DB.Query(queryEmail, u.Email)
	if err != nil {
		return err
	}
	// returns true if email already exists in the table, or false if not exists
	if rowsEmail.Next() {
		return errors.New("user with this email already exists")
	}

	return nil
}
