package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jj0308/GoTaskv2/handlers"
)

// SetupRouter creates and configures a new gin Engine
func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Instantiate handlers
	h := handlers.NewHandler(db)

	// Set up routes
	r.POST("/users/:id/events", h.AddUserToEvent)
	r.POST("/meetings", h.CreateMeetingAndInvitations)
	r.PUT("/invitations/:id", h.UpdateInvitationStatus)
	r.GET("/users/:id/invitations", h.GetAllInvitationsForUser)
	r.GET("/users/:id/meetings", h.GetAllMeetingsForUser)

	return r
}