package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"meeting/main/master/models"
	"strconv"
)

type MeetingRepoImpl struct {
	db *sql.DB
}

func (s MeetingRepoImpl) GetAllMeeting() ([]*models.Meeting, error) {
	dataMeeting := []*models.Meeting{}
	query := "select * from Meeting"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Meeting := models.Meeting{}
		var err = data.Scan(&Meeting.ID, &Meeting.TempatMeeting, &Meeting.Client, &Meeting.TanggalMeeting, &Meeting.WaktuMeeting)
		if err != nil {
			return nil, err
		}
		dataMeeting = append(dataMeeting, &Meeting)
	}
	fmt.Print("Endpoint")
	return dataMeeting, nil

}

func (s MeetingRepoImpl) GetMeetingByID(id string) (*models.Meeting, error) {
	query := "select * from Meeting WHERE idMeeting=?"
	result := s.db.QueryRow(query, id)

	var Meeting models.Meeting
	err := result.Scan(&Meeting.ID, &Meeting.TempatMeeting, &Meeting.Client, &Meeting.TanggalMeeting, &Meeting.WaktuMeeting)
	if err != nil {
		return nil, errors.New("User ID Not Found")
	}

	return &Meeting, nil
}

// HandlePOSTMeeting will POST a new Meeting data
func (s MeetingRepoImpl) HandlePOSTMeeting(d models.Meeting) (*models.Meeting, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmnt, _ := tx.Prepare(`INSERT INTO Meeting(tempat_meeting,client,tanggal_meeting,waktu) VALUES (?,?,?,?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.TempatMeeting, d.Client, d.TanggalMeeting, d.WaktuMeeting)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	fmt.Print("Endpoint")
	return s.GetMeetingByID(strconv.Itoa(int(lastInsertID)))
}

// HandleUPDATEMeeting is used for UPDATE data Meeting
func (s MeetingRepoImpl) HandleUPDATEMeeting(id string, data models.Meeting) (*models.Meeting, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE Meeting SET tempat_meeting = ?, client = ?,tanggal_meeting=?,waktu=? WHERE idMeeting=?`,
		data.TempatMeeting, data.Client, data.TanggalMeeting, data.WaktuMeeting, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	checkAvaibility, err := s.GetMeetingByID(id)
	if err != nil {
		return nil, err
	}
	fmt.Print("Endpoint")
	return checkAvaibility, nil
}

// HandleDELETEMeeting for DELETE single data from Meeting
func (s MeetingRepoImpl) HandleDELETEMeeting(id string) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return "Database Error", err
	}

	_, err = s.GetMeetingByID(id)
	if err != nil {
		return "", err
	}

	result, err := tx.Exec("DELETE FROM Meeting WHERE idMeeting=?", id)
	rowAffected, _ := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return strconv.Itoa(int(rowAffected)), err
	}
	tx.Commit()
	fmt.Print("Endpoint")
	return strconv.Itoa(int(rowAffected)) + " Affected", nil
}

func (s MeetingRepoImpl) GetAllMeetingByDate(day string) ([]*models.Meeting, error) {
	dataMeeting := []*models.Meeting{}
	query := "select * from Meeting where day(tanggal_meeting)=?"
	data, err := s.db.Query(query, day)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Meeting := models.Meeting{}
		var err = data.Scan(&Meeting.ID, &Meeting.TempatMeeting, &Meeting.Client, &Meeting.TanggalMeeting, &Meeting.WaktuMeeting)
		if err != nil {
			return nil, err
		}
		dataMeeting = append(dataMeeting, &Meeting)
	}
	fmt.Println("Endpoint")
	return dataMeeting, nil
}

func InitMeetingRepoImpl(db *sql.DB) MeetingRepository {
	return MeetingRepoImpl{db}
}
