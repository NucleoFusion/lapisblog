package routes

import (
	"database/sql"
	profileRoute "lapisblog/routes/profile"
	"net/http"
)

func GetRoutesMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/profile/", http.StripPrefix("/profile", profileRoute.InitMux(db)))

	return mux
}
