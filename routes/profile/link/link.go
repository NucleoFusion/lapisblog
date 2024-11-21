package link

import (
	"database/sql"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add", &addLink{Db: db})
	mux.Handle("GET /get", &GetLinkRoute{Db: db})
	mux.Handle("POST /remove/{id}", &RemoveLinkRoute{Db: db})

	return mux
}
