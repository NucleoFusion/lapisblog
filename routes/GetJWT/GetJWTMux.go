package getjwt

import (
	"database/sql"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /get", &GetJWT{Db: db})

	return mux
}
