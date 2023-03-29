package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jj0308/GoTaskv2/helpers"
	"github.com/jj0308/GoTaskv2/models"
)

type Handler struct {
	helper *helpers.Helper
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		helper: helpers.NewHelper(db),
	}
}

//user can participate in several events
// AddUserToEvent method
func (h *Handler) AddUserToEvent(c *gin.Context) {
		userID := c.Param("id")
		eventID := c.PostForm("event_id")

		if err := h.helper.ValidateUserAndEvent(userID, eventID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		participantExists, err := h.helper.CheckParticipantExistence(userID, eventID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking participant existence"})
			return
		}

		if participantExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is already a participant in the event"})
			return
		}

		if err := h.helper.AddParticipant(userID, eventID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding user to event"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User added to event successfully"})
}

//after creating a meeting ,invitations are created for every meeting invitee 
// CreateMeetingAndInvitations method
func (h *Handler) CreateMeetingAndInvitations(c *gin.Context) {
		var meetingRequest models.MeetingRequest

		if err := c.BindJSON(&meetingRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := h.helper.CheckParticipants(meetingRequest); err != nil {
			
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		meetingID, err := h.helper.CreateMeeting(meetingRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating meeting"})
			return
		}

		if err := h.helper.CreateInvitations(meetingRequest, meetingID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating invitation"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Meeting and invitations created successfully", "meeting_id": meetingID})
}

//invitations can be accepted or rejected
// UpdateInvitationStatus method
func (h *Handler) UpdateInvitationStatus(c *gin.Context) {
	invitationID := c.Param("id")

	if invitationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invitation ID is required"})
		return
	}

	var statusUpdate =  models.StatusUpdate

	if err := c.BindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if !h.helper.IsValidStatus(statusUpdate.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
		return
	}

	invitationExists, err := h.helper.CheckInvitationExists(invitationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking invitation existence"})
		return
	}

	if !invitationExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invitation does not exist"})
		return
	}


	// Update the invitation status and schedule the meeting if all invitees have accepted
	err = h.helper.UpdateInvitationStatusAndScheduleMeeting(invitationID, statusUpdate.Status)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the invitation status and scheduling the meeting"})
		log.Printf("Error updating the invitation status and scheduling the meeting: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invitation status updated successfully"})
}



//the status of each invitation must be tracked
//every user should have access to all invitations and scheduled meetings 
// GetAllInvitationsForUser method
func (h *Handler) GetAllInvitationsForUser(c *gin.Context){
		userID := c.Param("id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		userExists, err := h.helper.CheckUserExists(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user existence"})
			return
		}

		if !userExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
			return
		}

		invitations, err := h.helper.FetchUserInvitations(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching invitations"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"invitations": invitations})
}

// GetAllMeetingsForUser method
func (h *Handler) GetAllMeetingsForUser(c *gin.Context) {
	userID := c.Param("id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		userExists, err := h.helper.CheckUserExists(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking user existence"})
			return
		}

		if !userExists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
			return
		}

		meetings, err := h.helper.FetchUserMeetings(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching meetings"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"meetings": meetings})
}
