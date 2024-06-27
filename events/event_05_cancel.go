package events

import (
	"events-app/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Cancel(ctx *gin.Context) {
	// parse event id from the http request header
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id", "reason": err.Error()})
		return
	}

	// check if the event exists in the events table and retrieve the event if exists
	event, err := GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event does not exist", "reason": err.Error()})
		return
	}

	// check if user requesting event cancellation owns the requested event and can change it
	authUserID := ctx.GetInt64("userID")
	if authUserID != event.UserID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	err = event.cancel(eventID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Event cancellation failed", "reason": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "Event successfully cancelled!"})
}

func (e *Event) cancel(eventID int64) error {
	query := `UPDATE events SET status = $1 WHERE id = $2;`
	_, err := db.DB.Exec(query, "false", eventID)
	if err != nil {
		return err
	}

	return nil
}
