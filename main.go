package main

import (
	"fmt"
	"lapisblog/routes"
	"net/http"
)

//TODO Profile Addition routes
//TODO One-Time API Key Authentication
//TODO JWT Auth Routes

func main() {
	mux := routes.GetRoutesMux()

	fmt.Println("Listening at 5555.")
	http.ListenAndServe(":5555", mux)

	// payload := auth.CreatePayload(1, "User")

	// jwt := payload.CreateJWT()

	// st, err := jwt.EncodeJWT()
	// if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	// fmt.Println(st)

}
