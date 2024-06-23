package events

import "time"

type EventNew struct {
	ID          int64
	Name        string    `json:"name" binding:"required,min=3"`
	Description string    `json:"description" binding:"required,min=10"`
	Location    string    `json:"location" binding:"required,min=3"`
	Date        time.Time `json:"date" binding:"required" time_format:"2006-12-31"` // time_format is ignored by Gin
	UserID      int64
}

// this type is used with GET http request for retrieving users events and showing them to the user
type Event struct {
	Name        string
	Description string
	Location    string
	Date        time.Time
	UserID      int64
}

/* Regarding time_format validation, see below explanation why it does not work with JSON.
Figured out the answer. It has nothing to do with Gin, but what the golang JSON Decoder expects for time values.
It must be in the format of RFC3339 in order to be successfully decoded when using the JSON binder for ShouldBindJSON, or BindJSON.
However, it should be addressed that this creates an inconsistency when binding from form data vs binding from JSON data.
When binding form data, if you set the time_format tag and try to send anything other than the time_format you specify,
it will fail, but with JSON it totally ignores the time_format and throws an error if it is not in RFC3339 format.
If not fixed in Gin, this should at least be addressed in the docs. I'd be happy to make a PR!
https://github.com/gin-gonic/gin/issues/1193#issuecomment-350498604

---- SOLUTION ----
type myTime time.Time

var _ json.Unmarshaler = &myTime{}

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = myTime(t)
	return nil
}

type Class struct {
	StartAt     myTime `json:"start_at" binding:"required"`
	ChallengeID uint   `json:"challenge_id" gorm:"index" binding:"required"`
}
*/
