package link

import (
	"database/sql"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add", &addLink{Db: db})

	return mux
}
