package routes

import (
	profileRoute "lapisblog/routes/profile"
	"net/http"
)

func GetRoutesMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/profile/", http.StripPrefix("/profile", profileRoute.InitMux()))

	return mux
}
