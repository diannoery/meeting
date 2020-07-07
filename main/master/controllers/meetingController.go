package controllers

import (
	"encoding/json"
	"meeting/main/master/models"
	"meeting/main/master/usecases"
	"meeting/utils/message"
	"meeting/utils/tools"
	"net/http"

	"github.com/gorilla/mux"
)

type MeetingHandler struct {
	MeetingUseCase usecases.MeetingUseCase
}

func MeetingController(r *mux.Router, service usecases.MeetingUseCase) {
	MeetingHandler := MeetingHandler{service}
	r.HandleFunc("/meeting", MeetingHandler.ListMeeting).Methods(http.MethodGet)
	r.HandleFunc("/meeting/{id}", MeetingHandler.MeetingByID).Methods(http.MethodGet)
	r.HandleFunc("/meeting", MeetingHandler.HandlePostMeeting).Methods(http.MethodPost)
	r.HandleFunc("/meeting/{id}", MeetingHandler.HandlePutMeeting).Methods(http.MethodPut)
	r.HandleFunc("/meeting/{id}", MeetingHandler.HandleDeleteMeeting).Methods(http.MethodDelete)
	r.HandleFunc("/jadwal/{day}", MeetingHandler.JadwalMeeting).Methods(http.MethodGet)

}

func (s MeetingHandler) ListMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	Meetings, err := s.MeetingUseCase.GetMeeting()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.Respone("Transaction Failed", http.StatusBadRequest, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(message.Respone("Transaction Success", http.StatusOK, Meetings))
}

func (s MeetingHandler) MeetingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	Meeting, err := s.MeetingUseCase.GetMeetingByID(tools.GetPathVar("id", r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.Respone("Search Failed", http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message.Respone("Search Success", http.StatusOK, Meeting))
}

func (s MeetingHandler) HandlePostMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var parseMeeting models.Meeting
	tools.Parser(r, &parseMeeting)

	Meeting, err := s.MeetingUseCase.HandlePOSTMeeting(parseMeeting)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.Respone("Transaction Failed", http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message.Respone("Transaction Success", http.StatusOK, Meeting))
}

func (s MeetingHandler) HandlePutMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var getMeeting models.Meeting
	tools.Parser(r, &getMeeting)

	Meeting, err := s.MeetingUseCase.HandleUPDATEMeeting(tools.GetPathVar("id", r), getMeeting)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.Respone("Update Failed", http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message.Respone("Update Success", http.StatusOK, Meeting))
}

func (s MeetingHandler) HandleDeleteMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	Meeting, err := s.MeetingUseCase.HandleDELETEMeeting(tools.GetPathVar("id", r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.Respone("Delete Meeting By ID Failed", http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message.Respone("Delete Meeting By ID Success", http.StatusOK, Meeting))
}

func (s MeetingHandler) JadwalMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	Meeting, err := s.MeetingUseCase.GetMeetingByDate(tools.GetPathVar("day", r))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.Respone("Transaction Failed", http.StatusBadRequest, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(message.Respone("Meeting BY Date", http.StatusOK, Meeting))
}

	