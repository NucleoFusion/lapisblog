package profileRoute

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"lapisblog/statics"
	"net/http"
)

type updateProfile struct {
	Db *sql.DB
}

func (s *updateProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}

	body := r.PostForm

	profile, err := DecodeBody(&body)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	data, _ := json.Marshal(profile)

	// checkExists := make(chan bool, 1)

	// go InsertIntoProfile(profile, s.Db, checkExists)

	// exists := <-checkExists

	// if exists {
	// io.WriteString(w, "user already exists")
	// return
	// }

	io.Writer.Write(w, data)
}

func UpdateIntoProfile(profile *statics.Profile, db *sql.DB, exists chan bool) {
	var (
		name        = ReturnNULL(profile.Name)
		email       = profile.Email
		linkedin    = ReturnNULL(profile.LinkedIn)
		description = ReturnNULL(profile.Description)
		birthDate   = profile.BirthDate
	)

	res := db.QueryRow("SELECT * FROM profile WHERE email = $1", email)
	if res.Scan() == sql.ErrNoRows {
		fmt.Println("Doc Found")
		exists <- true
		return
	}

	exists <- false

	_, err := db.Exec("INSERT INTO profile (name, email, linkedin, description, birthdate) VALUES ($1, $2, $3, $4, $5)", name, email, linkedin, description, birthDate)
	if err != nil {
		fmt.Println(err.Error())
	}
}
