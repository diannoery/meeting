package models

type Meeting struct {
	ID             string `json:"id"`
	TempatMeeting  string `json:"TempatMeeting"`
	Client         string `json:"Client"`
	TanggalMeeting string `json:"Tanggal"`
	WaktuMeeting   string `json:"Waktu"`
}
