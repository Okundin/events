package events

import (
	"events-app/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvent(ctx *gin.Context) {
	//parse GET http request for event id
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check if event exists in the event table or retrieve it
	event, err := GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Event not found!"})
		return
	}

	// check if the event belongs to the user
	authUserID := ctx.GetInt64("userID")
	if authUserID != event.UserID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized!"})
		return
	}

	ctx.JSON(http.StatusOK, &event)
}

func GetEventByID(eventID int64) (*Event, error) {
	query := `SELECT name, description, location, date, userid FROM events WHERE id = $1;`
	row := db.DB.QueryRow(query, eventID)

	// unpacking the retrieved row
	var event Event
	err := row.Scan(&event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
