package models

type Organization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
}

type Event struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

type EventParticipant struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}

type Meeting struct {
	ID          int    `json:"id"`
	EventID     int    `json:"event_id"`
	OrganizerID int    `json:"organizer_id"`
	DateTime    string `json:"datetime"`
}

type Invitation struct {
	ID         int    `json:"id"`
	MeetingID  int    `json:"meeting_id"`
	InviteeID  int    `json:"invitee_id"`
	Status     string `json:"status"`
}
type MeetingRequest struct {
	EventID     int    `json:"event_id"`
	OrganizerID int    `json:"organizer_id"`
	DateTime    string `json:"datetime"`
	InviteeIDs  []int  `json:"invitee_ids"`
}

var StatusUpdate struct {
	Status string `json:"status"`
}