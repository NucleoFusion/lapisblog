package profileRoute

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lapisblog/auth"
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

	checkExists := make(chan bool, 1)
	keyString := make(chan string, 1)
	giveKey := make(chan string, 1)

	go InsertIntoProfile(profile, s.Db, checkExists, giveKey)
	go KeyBuffer(profile, keyString)

	exists := <-checkExists
	if exists {
		io.WriteString(w, "user already exists")
		return
	}

	key := <-keyString
	giveKey <- key
	profile.Key = key

	data, _ := json.Marshal(profile)
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

	if prof.Email == "" {
		return &prof, errors.New("email not provided")
	}

	return &prof, nil
}

func InsertIntoProfile(profile *statics.Profile, db *sql.DB, exists chan bool, keyString chan string) {
	var (
		name        = ReturnNULL(profile.Name)
		email       = profile.Email
		linkedin    = ReturnNULL(profile.LinkedIn)
		description = ReturnNULL(profile.Description)
		birthDate   = profile.BirthDate
	)

	fmt.Println(profile)

	res := db.QueryRow("SELECT * FROM profile WHERE email = $1", email)
	if res.Scan() != sql.ErrNoRows {
		exists <- true
		return
	}

	exists <- false
	key := <-keyString

	_, err := db.Exec("INSERT INTO profile (name, email, linkedin, description, birthdate, key) VALUES ($1, $2, $3, $4, $5, $6)", name, email, linkedin, description, birthDate, key)
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

func KeyBuffer(profile *statics.Profile, keyChan chan string) {
	keyChan <- auth.GetKey(profile)
}
