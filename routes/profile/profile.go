package profileRoute

import (
	"database/sql"
	"lapisblog/routes/profile/education"
	"lapisblog/routes/profile/link"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add", &addProfileStruct{Db: db})
	mux.Handle("POST /update", &updateProfile{Db: db})

	mux.Handle("/link/", http.StripPrefix("/link", link.InitMux(db)))
	mux.Handle("/education/", http.StripPrefix("/education", education.InitMux(db)))

	return mux
}
