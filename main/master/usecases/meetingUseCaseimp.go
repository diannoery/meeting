package usecases

import (
	"meeting/main/master/models"
	"meeting/main/master/repositories"
	"meeting/utils/tools"
)

type MeetingUsecaseImpl struct {
	MeetingRepo repositories.MeetingRepository
}

func (m MeetingUsecaseImpl) GetMeeting() ([]*models.Meeting, error) {
	Meeting, err := m.MeetingRepo.GetAllMeeting()
	if err != nil {
		return nil, err
	}

	return Meeting, nil
}

func (m MeetingUsecaseImpl) GetMeetingByID(id string) (*models.Meeting, error) {
	Meeting, err := m.MeetingRepo.GetMeetingByID(id)
	if err != nil {
		return nil, err
	}
	return Meeting, nil
}

func (m MeetingUsecaseImpl) HandlePOSTMeeting(d models.Meeting) (*models.Meeting, error) {

	err := tools.ValidateInputNotNil(d.TempatMeeting, d.Client, d.TanggalMeeting, d.WaktuMeeting)
	if err != nil {
		return &d, err
	}
	result, err := m.MeetingRepo.HandlePOSTMeeting(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m MeetingUsecaseImpl) HandleUPDATEMeeting(id string, data models.Meeting) (*models.Meeting, error) {
	err := tools.ValidateInputNotNil(data.TempatMeeting, data.Client, data.TanggalMeeting, data.WaktuMeeting)
	if err != nil {
		return &data, err
	}
	result, err := m.MeetingRepo.HandleUPDATEMeeting(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MeetingUsecaseImpl) HandleDELETEMeeting(id string) (string, error) {
	result, err := m.MeetingRepo.HandleDELETEMeeting(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m MeetingUsecaseImpl) GetMeetingByDate(day string) ([]*models.Meeting, error) {
	Meeting, err := m.MeetingRepo.GetAllMeetingByDate(day)

	if err != nil {
		return nil, err
	}

	return Meeting, nil
}

func InitMeetingUseCase(MeetingRepo repositories.MeetingRepository) MeetingUseCase {
	return MeetingUsecaseImpl{MeetingRepo}
}
