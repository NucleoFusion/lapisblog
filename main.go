package main

import (
	"fmt"
	"lapisblog/auth/jwt"
)

func main() {
	// mux := routes.GetRoutesMux()

	// fmt.Println("Listening at 5555.")
	// http.ListenAndServe(":5555", mux)

	k, _ := jwt.GetKey()

	msg, err := jwt.EncryptMsg([]byte("Hello"), k)

	fmt.Println(string(msg), err)

	decmsg, _ := jwt.DecryptMsg(msg, k)

	fmt.Println(string(decmsg))
}
