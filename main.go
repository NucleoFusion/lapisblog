package main

import (
	"context"
	"fmt"
	"lapisblog/database"
)

//TODO Profile Addition routes
//TODO One-Time API Key Authentication
//TODO JWT Auth Routes

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

	db, err := db.ConnectToDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err := db.QueryContext(context.Background(), "SELECT * FROM users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var (
		id       int
		username string
		email    string
		password string
	)
	for result.Next() {
		result.Scan(&id, &username, &email, &password)
		fmt.Println(id, username, email, password)
	}
}
