package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jj0308/GoTaskv2/models"
	"github.com/jj0308/GoTaskv2/storage"
)

type Helper struct {
	db *sql.DB
}

func NewHelper(db *sql.DB) *Helper {
	return &Helper{
		db: db,
	}
}

func (h *Helper) AddParticipant(userID, eventID string) error {
	_, err := h.db.Exec(storage.AddParticipantQuery, sql.Named("user_id", userID), sql.Named("event_id", eventID))

	return err
}
func (h *Helper) CheckParticipants(meetingRequest models.MeetingRequest) error {
	userIDs := append([]int{meetingRequest.OrganizerID}, meetingRequest.InviteeIDs...)

	for _, userID := range userIDs {
		participantExists, err := h.IsParticipant(userID, meetingRequest.EventID)
		if err != nil {
			return err
		}

		if !participantExists {
			return errors.New("Organizer or Invitee does not belong to the specified event")
		}
	}

	return nil
}

func (h *Helper) CreateInvitations(meetingRequest models.MeetingRequest, meetingID int64) error {
		for _, inviteeID := range meetingRequest.InviteeIDs {
			_, err := h.db.Exec(storage.CreateInvitationQuery, sql.Named("meeting_id", meetingID), sql.Named("invitee_id", inviteeID), sql.Named("status", "pending"))
			if err != nil {
					return err
			}
		}

		return nil
}

func (h *Helper) CreateMeeting(meetingRequest models.MeetingRequest) (int64, error) {
	var meetingID int64
	err := h.db.QueryRow(storage.CreateMeetingQuery,
		sql.Named("event_id", meetingRequest.EventID),
		sql.Named("organizer_id", meetingRequest.OrganizerID),
		sql.Named("datetime", meetingRequest.DateTime)).Scan(&meetingID)

	return meetingID, err
}

func (h *Helper) UpdateInvitationStatus(invitationID, status string) error {
	_, err := h.db.Exec(storage.UpdateInvitationStatusQuery, sql.Named("status", status),
	sql.Named("invitationID", invitationID))

	return err
}

func (h *Helper) FetchUserInvitations(userID string) ([]models.Invitation, error) {
	rows, err := h.db.Query(storage.FetchUserInvitationsQuery, sql.Named("inviteeID",userID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invitations := []models.Invitation{}
	for rows.Next() {
		var invitation models.Invitation
		err := rows.Scan(&invitation.ID, &invitation.MeetingID, &invitation.InviteeID, &invitation.Status)
		if err != nil {
			return nil, err
		}
		invitations = append(invitations, invitation)
	}

	return invitations, nil
}


func (h *Helper) CheckEntityExistence(table, id string) (bool, error) {
	var exists bool
	query := fmt.Sprintf(storage.CheckEntityExistenceQuery, table)
	err := h.db.QueryRow(query, sql.Named("id", id)).Scan(&exists)
	if err != nil {
		log.Printf("Error checking %s existence: %s", table, err.Error())
	}
	return exists, err
}

func (h *Helper) CheckParticipantExistence(userID, eventID string) (bool, error) {
	var participantExists bool
	err := h.db.QueryRow(storage.CheckParticipantExistenceQuery, sql.Named("userID", userID), sql.Named("eventID", eventID)).Scan(&participantExists)
	return participantExists, err
}

func (h *Helper) IsParticipant(userID int, eventID int) (bool, error) {
	var participantExists bool
	err := h.db.QueryRow(storage.CheckUserParticipantQuery,
		sql.Named("userID", userID),
		sql.Named("eventID", eventID)).Scan(&participantExists)

	return participantExists, err
}

func (h *Helper) CheckInvitationExists(invitationID string) (bool, error) {
	var invitationExists bool
	err := h.db.QueryRow(storage.CheckInvitationExistsQuery,
		sql.Named("invitationID", invitationID)).Scan(&invitationExists)

	return invitationExists, err
}
func (h *Helper) CheckUserExists(userID string) (bool, error) {
	var userExists bool
	err := h.db.QueryRow(storage.CheckUserExistsQuery, sql.Named("userID",userID)).Scan(&userExists)
	return userExists, err
}

func (h *Helper) ValidateUserAndEvent(userID, eventID string) error {
	if userID == "" || eventID == "" {
		return errors.New("User ID and Event ID are required")
	}

	userExists, err := h.CheckEntityExistence("users", userID)
	if err != nil {
		return errors.New("Error checking user existence")
	}

	eventExists, err := h.CheckEntityExistence("events", eventID)
	if err != nil {
		return errors.New("Error checking event existence")
	}

	if !userExists || !eventExists {
		return errors.New("User or Event does not exist")
	}

	return nil
}
func (h *Helper) IsValidStatus(status string) bool {
	return status == "accepted" || status == "rejected"
}
func (h *Helper) FetchUserMeetings(userID string) ([]models.Meeting, error) {
	rows, err := h.db.Query(storage.FetchUserMeetingsQuery, sql.Named("organizerID", userID), sql.Named("inviteeID", userID))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meetings := []models.Meeting{}
	for rows.Next() {
		var meeting models.Meeting
		err := rows.Scan(&meeting.ID, &meeting.EventID, &meeting.OrganizerID, &meeting.DateTime)
		if err != nil {
		return nil, err
		}
		meetings = append(meetings, meeting)
		}
		return meetings, nil
	}

	func (h *Helper) UpdateInvitationStatusAndScheduleMeeting(invitationID string, status string) error {
    // First, update the invitation status
    err := h.UpdateInvitationStatus(invitationID, status)
    if err != nil {
        return fmt.Errorf("error updating invitation status: %v", err)
    }

    // If the status is not "accepted", there's no need to check for scheduling the meeting
    if status != "accepted" {
        return nil
    }

    meetingID, err := h.GetMeetingIDByInvitationID(invitationID)
    if err != nil {
			return err
    }

    allAccepted, err := h.CheckIfAllInviteesAccepted(meetingID)
    if err != nil {
			return err
		}

		// If all invitees have accepted the invitation, schedule the meeting
		if allAccepted {
			err = h.ScheduleMeeting(meetingID)
			if err != nil {
			 return err
			}
		}

		return nil
}
	
	func (h *Helper) GetMeetingIDByInvitationID(invitationID string) (int, error) {
		var meetingID int
	
		
		err := h.db.QueryRow(storage.GetMeetingIDByInvitationIDQuery,sql.Named("invitationID", invitationID)).Scan(&meetingID)
		if err != nil {
			if err == sql.ErrNoRows {
				return 0, fmt.Errorf("no invitation found with ID: %s", invitationID)
			}
			return 0, err
		}
	
		return meetingID, nil
	}

	func (h *Helper) CheckIfAllInviteesAccepted(meetingID int) (bool, error) {
		var totalInvitees, acceptedInvitees int
	
		
		err := h.db.QueryRow(storage.TotalQuery,sql.Named("meetingID", meetingID)).Scan(&totalInvitees)
		if err != nil {
			return false, fmt.Errorf("error querying total invitees: %v", err)
		}
	
		
		err = h.db.QueryRow(storage.AcceptedQuery,sql.Named("meetingID", meetingID)).Scan(&acceptedInvitees)
		if err != nil {
			return false, fmt.Errorf("error querying accepted invitees: %v", err)
		}
	
		return totalInvitees == acceptedInvitees, nil
	}

	func (h *Helper) ScheduleMeeting(meetingID int) error {
		
		result, err := h.db.Exec(storage.ScheduleMeetingQuery, sql.Named("meetingID", meetingID))
		if err != nil {
			return err
		}
	
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
	
		if rowsAffected == 0 {
			return err
		}
	
		return nil
	}