package profileRoute

import "net/http"

func InitMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/add", &addProfileStruct{})

	return mux
}
