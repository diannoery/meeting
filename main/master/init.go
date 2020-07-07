package master

import (
	"database/sql"
	"fmt"
	"meeting/main/master/controllers"
	"meeting/main/master/repositories"
	"meeting/main/master/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {

	MeetingRepo := repositories.InitMeetingRepoImpl(db)
	MeetingUseCase := usecases.InitMeetingUseCase(MeetingRepo)
	controllers.MeetingController(r, MeetingUseCase)

	r.NotFoundHandler = http.HandlerFunc(notFound)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>404 Status Not Found</h1>`)
}
