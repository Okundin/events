package events

import (
	"events-app/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	// retrieving all events created by the user from the DB
	query := `SELECT name, description, location, date, userid FROM events WHERE userid = $1;`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// unpacking the retrieved rows to the events slice
	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.Name, &e.Description, &e.Location, &e.Date, &e.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		events = append(events, e)
	}

	ctx.JSON(http.StatusOK, &events)
}
