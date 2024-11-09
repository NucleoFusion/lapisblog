package profileRoute

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lapisblog/statics"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type addProfileStruct struct {
	Db *sql.DB
}

func (s *addProfileStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	go InsertIntoProfile(profile, s.Db)

	io.Writer.Write(w, data)
}

func DecodeBody(body *url.Values) (*statics.Profile, error) {
	prof := statics.Profile{}

	for key, val := range *body {
		switch key {
		case "email":
			prof.Email = val[0]
		case "name":
			prof.Name = val[0]
		case "description":
			prof.Description = val[0]
		case "linkedin":
			prof.LinkedIn = val[0]
		case "birthDate":
			dates := strings.Split(val[0], "/")
			if len(dates) != 3 {
				return &prof, errors.New("invalid date, the format is dd/mm/yyyy")
			}

			day, err := strconv.Atoi(dates[2])
			month, err2 := strconv.Atoi(dates[1])
			year, err3 := strconv.Atoi(dates[0])
			if err != nil || err2 != nil || err3 != nil {
				return &prof, errors.New("invalid date, the format is dd/mm/yyyy")
			}

			date := time.Date(day, time.Month(month), year, 0, 0, 0, 0, time.Local).Unix()

			prof.BirthDate = date
		}
	}

	return &prof, nil
}

func InsertIntoProfile(profile *statics.Profile, db *sql.DB) {
	var (
		name        = ReturnNULL(profile.Name)
		email       = ReturnNULL(profile.Email)
		linkedin    = ReturnNULL(profile.LinkedIn)
		description = ReturnNULL(profile.Description)
		birthDate   = profile.BirthDate
	)

	_, err := db.Exec("INSERT INTO profile (name, email, linkedin, description, birthdate) VALUES ($1, $2, $3, $4, $5)", name, email, linkedin, description, birthDate)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ReturnNULL(s string) string {
	if s == "" {
		return "NULL"
	}

	return s
}
