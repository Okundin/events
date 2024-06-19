package events

import (
	"events-app/db"
	"events-app/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(ctx *gin.Context) {
	//parsing the POST http request
	var err error
	var event EventNew
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// getting userid from the user token & assigning userid to the event struct
	userID := ctx.GetInt64("userID")
	event.UserID = userID

	// verifying if user exists in the DB
	_, err = users.GetUserByID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found!"})
		return
	}

	err = event.create()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"1 Congrats!":   "New event successfully created",
		"2 Name":        event.Name,
		"3 Description": event.Description,
		"4 Location":    event.Location,
		"Date":          event.Date,
	})
}

// create method writes event to the event table of the DB
func (e *EventNew) create() error {
	query := `INSERT INTO events (name, description, location, date, userid) VALUES ($1, $2, $3, $4, $5);`
	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.Date, e.UserID)
	if err != nil {
		return err
	}

	return nil
}
