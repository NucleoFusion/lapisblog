package routes

import (
	"database/sql"
	getjwt "lapisblog/routes/GetJWT"
	profileRoute "lapisblog/routes/profile"
	"net/http"
)

func GetRoutesMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/profile/", http.StripPrefix("/profile", profileRoute.InitMux(db)))
	mux.Handle("/jwt/", http.StripPrefix("/jwt", getjwt.InitMux(db)))

	return mux
}
