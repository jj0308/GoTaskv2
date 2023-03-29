package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jj0308/GoTaskv2/handlers"
)

// SetupRouter creates and configures a new gin Engine
//SetupRouter that takes a pointer to a sql.DB object as an argument and returns a pointer to a gin.Engine object.
func SetupRouter(db *sql.DB) *gin.Engine {
	// This line initializes a new Gin router instance with the default settings, including the default middleware (like the logger and recovery middleware).
	r := gin.Default()
	// Instantiate handlers
	// This instance will be used to handle the different routes defined in the router.
	h := handlers.NewHandler(db)

	r.POST("/users/:id/events", h.AddUserToEvent)
	r.POST("/meetings", h.CreateMeetingAndInvitations)
	r.PUT("/invitations/:id", h.UpdateInvitationStatus)
	r.GET("/users/:id/invitations", h.GetAllInvitationsForUser)
	r.GET("/users/:id/meetings", h.GetAllMeetingsForUser)

	return r
}