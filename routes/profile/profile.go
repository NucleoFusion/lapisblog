package profileRoute

import (
	"database/sql"
	"lapisblog/routes/profile/education"
	"lapisblog/routes/profile/follow"
	"lapisblog/routes/profile/link"
	"lapisblog/routes/profile/tags"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add", &addProfileStruct{Db: db})
	mux.Handle("POST /update", &updateProfile{Db: db})

	mux.Handle("/link/", http.StripPrefix("/link", link.InitMux(db)))
	mux.Handle("/education/", http.StripPrefix("/education", education.InitMux(db)))
	mux.Handle("/tags/", http.StripPrefix("/tags", tags.InitMux(db)))
	mux.Handle("/follow/", http.StripPrefix("/follow", follow.InitMux(db)))

	return mux
}
