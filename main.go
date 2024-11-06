package main

import (
	"fmt"
	"lapisblog/auth"
)

func main() {
	// mux := routes.GetRoutesMux()

	// fmt.Println("Listening at 5555.")
	// http.ListenAndServe(":5555", mux)

	payload, err := auth.CreatePayload(1, "Admin", -1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	JWT := auth.GetJWT(payload)

	encrypted, err := auth.EncryptJWT(JWT)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("ENCRYPTED", encrypted)
}
