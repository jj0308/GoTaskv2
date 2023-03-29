package storage

const (
	CheckEntityExistenceQuery = "SELECT IIF(EXISTS(SELECT 1 FROM %s WHERE id = @id), 1, 0)"

	CheckParticipantExistenceQuery = "SELECT IIF(EXISTS(SELECT 1 FROM event_participants WHERE user_id = @userID AND event_id = @eventID), 1, 0)"

	AddParticipantQuery = "INSERT INTO event_participants (user_id, event_id) VALUES (@user_id, @event_id)"

	CheckUserParticipantQuery = "SELECT IIF(EXISTS(SELECT 1 FROM event_participants WHERE user_id = @userID AND event_id = @eventID), 1, 0)"

	CreateMeetingQuery = "INSERT INTO meetings (event_id, organizer_id, datetime) OUTPUT INSERTED.id VALUES (@event_id, @organizer_id, @datetime)"

	CreateInvitationQuery = "INSERT INTO invitations (meeting_id, invitee_id, status) VALUES (@meeting_id, @invitee_id, @status)"
	CheckInvitationExistsQuery ="SELECT IIF(EXISTS(SELECT 1 FROM invitations WHERE id = @invitationID), 1, 0)"
	UpdateInvitationStatusQuery ="UPDATE invitations SET status = @status WHERE id = @invitationID"

	FetchUserInvitationsQuery = "SELECT id, meeting_id, invitee_id, status FROM invitations WHERE invitee_id = @inviteeID"

	CheckUserExistsQuery = "SELECT IIF(EXISTS(SELECT 1 FROM users WHERE id = @userID), 1, 0)"

	FetchUserMeetingsQuery = `
	SELECT DISTINCT
		m.id, m.event_id, m.organizer_id, m.datetime
	FROM meetings m
	LEFT JOIN invitations i ON m.id = i.meeting_id
	WHERE m.organizer_id = @organizerID OR i.invitee_id = @inviteeID`

	ScheduleMeetingQuery = "UPDATE meetings SET scheduled = 1 WHERE id = @meetingID"
	GetMeetingIDByInvitationIDQuery = "SELECT meeting_id FROM invitations WHERE id = @invitationID"
	TotalQuery = "SELECT COUNT(*) FROM invitations WHERE meeting_id = @meetingID"
	AcceptedQuery = "SELECT COUNT(*) FROM invitations WHERE meeting_id = @meetingID AND status = 'accepted'"
)