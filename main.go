package main

import (
	"fmt"
	"lapisblog/auth"
)

func main() {
	// mux := routes.GetRoutesMux()

	// fmt.Println("Listening at 5555.")
	// http.ListenAndServe(":5555", mux)

	// payload := auth.CreatePayload(1, "User")

	// jwt := payload.CreateJWT()

	// st, err := jwt.EncodeJWT()
	// if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	// fmt.Println(st)

	st := "eyJhbGciOiJTSEEyNTYiLCJ0eXAiOiJKV1MiLCJjdHkiOiJKU09OIn0.eyJJZCI6MSwiUm9sZSI6IlVzZXIiLCJDcmVhdGVkQXQiOjE3MzEwNjY5ODcsIlZhbGlkRm9yIjowfQ.TBpANxlHZp8dcf1LyKXhLeKL4OmlwVBtFxLc2qiHu3E"

	s, err := auth.DecodeJWT(st)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(s)
}
