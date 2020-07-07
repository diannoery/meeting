package usecases

import "meeting/main/master/models"

type MeetingUseCase interface {
	GetMeeting() ([]*models.Meeting, error)
	GetMeetingByID(id string) (*models.Meeting, error)
	HandlePOSTMeeting(d models.Meeting) (*models.Meeting, error)
	HandleUPDATEMeeting(id string, data models.Meeting) (*models.Meeting, error)
	HandleDELETEMeeting(id string) (string, error)
	GetMeetingByDate(day string) ([]*models.Meeting, error)
}
