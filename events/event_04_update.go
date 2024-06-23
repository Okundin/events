package events

import (
	"events-app/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	// parse PUT http request
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request", "message": err.Error()})
		return
	}

	// check if event exists in the event table
	event, err := GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found!"})
		return
	}

	// comapre user id from the http request with the user id from the event to make sure if user is authorized to update the event
	authUserID := ctx.GetInt64("userID")
	if authUserID != event.UserID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized!"})
		return
	}

	// parsing JSON http request body elements to update/rewrite the event with new values
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse JSON body", "message": err.Error()})
		return
	}

	err = event.update(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Congrats": "Event successfully updated!"})
}

// update method updates event in the event table with new values from http request body section submitted by the user
func (e *Event) update(eventID int64) error {
	query := `UPDATE events SET name = $1, description = $2, location = $3, date = $4 WHERE id = $5;`
	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.Date, eventID)
	if err != nil {
		return err
	}

	return nil
}
