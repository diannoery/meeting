package repositories

import (
	"meeting/main/master/models"
)

type MeetingRepository interface {
	GetAllMeeting() ([]*models.Meeting, error)
	GetMeetingByID(id string) (*models.Meeting, error)
	HandlePOSTMeeting(d models.Meeting) (*models.Meeting, error)
	HandleUPDATEMeeting(id string, data models.Meeting) (*models.Meeting, error)
	HandleDELETEMeeting(id string) (string, error)
	GetAllMeetingByDate(day string) ([]*models.Meeting, error)
}
