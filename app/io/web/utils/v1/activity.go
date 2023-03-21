package v1

import (
	"github.com/dimonrus/gorest"
	"net/http"
	"tutorial/app/base"
	"tutorial/app/core"
	"tutorial/app/io/db/models"
)

// ResponseCurrentActivityHandler Utils get activity response
//
// swagger:response ResponseCurrentActivityHandler
type ResponseCurrentActivityHandler struct {
	// In: body
	Body struct {
		// Message
		// Required: true
		// Example: Memory usage
		Message string `json:"message,omitempty"`
		// Build info
		// Required: true
		Data models.ActivityList `json:"data"`
	}
}

// swagger:route GET utils/v1/activity Utils CurrentActivityHandler
//
// Utils. Current database activity
//
// Show database activity
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: ResponseCurrentTimeHandler
func CurrentActivityHandler(w http.ResponseWriter, r *http.Request) {
	activities, e := core.GetDatabaseActivity(base.App.GetDB())
	if e != nil {
		gorest.Send(w, gorest.NewErrorJsonResponse(e))
		return
	}
	gorest.Send(w, gorest.NewOkJsonResponse("All user activities", activities, nil))
}
