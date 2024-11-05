package main

import (
	"fmt"
	"lapisblog/routes"
	"net/http"
)

func main() {
	mux := routes.GetRoutesMux()

	fmt.Println("Listening at 5555.")
	http.ListenAndServe(":5555", mux)
}
