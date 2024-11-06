package main

import (
	"fmt"
	"lapisblog/auth"
)

func main() {
	// mux := routes.GetRoutesMux()

	// fmt.Println("Listening at 5555.")
	// http.ListenAndServe(":5555", mux)

	payload := auth.CreatePayload(1, "User")

	jwt := payload.CreateJWT()

	s, err := jwt.EncodeJWT()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(s)
}
