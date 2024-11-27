package education

import (
	"database/sql"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add", &AddEducation{Db: db})
	mux.Handle("POST /remove/{id}", &RemoveEducationRoute{Db: db})
	mux.Handle("GET /get", &GetEducationRoute{Db: db})

	return mux
}
