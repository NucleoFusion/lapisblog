package profileRoute

import "net/http"

func InitMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add", &addProfileStruct{})

	return mux
}
