package profileRoute

import (
	"database/sql"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add", &addProfileStruct{Db: db})
	mux.Handle("POST /update", &updateProfile{Db: db})

	return mux
}
