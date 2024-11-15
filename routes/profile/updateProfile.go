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

	key := r.URL.Query().Get("key")

	profile, err := DecodeBody(&body)
	if err != nil {
		if err.Error() != "email not provided" {
			io.WriteString(w, err.Error())
			return
		}
	}

	profile.Key = key

	checkExists := make(chan bool, 1)
	profChan := make(chan *statics.Profile, 1)
	sendMerged := make(chan *statics.Profile, 1)

	go UpdateIntoProfile(profile, checkExists, profChan, sendMerged, s.Db)

	exists := <-checkExists
	if !exists {
		io.WriteString(w, "user doesnt exists")
		return
	}

	prof := <-profChan
	MergeProfiles(profile, prof)
	sendMerged <- prof

	data, _ := json.Marshal(prof)

	io.Writer.Write(w, data)
}

func UpdateIntoProfile(profile *statics.Profile, exists chan bool, profileChan chan *statics.Profile, getMerged chan *statics.Profile, db *sql.DB) {
	res := db.QueryRow("SELECT * FROM profile WHERE key = $1", profile.Key)
	prof := statics.Profile{}

	var id int

	err := res.Scan(&id, &prof.Name, &prof.Email, &prof.LinkedIn, &prof.Description, &prof.BirthDate, &prof.Key)
	if err != nil {
		fmt.Println(err.Error())
		exists <- false
		return
	}

	exists <- true
	profileChan <- &prof

	completeProfile := <-getMerged

	var (
		name        = completeProfile.Name
		linkedin    = completeProfile.LinkedIn
		description = completeProfile.Description
		BirthDate   = completeProfile.BirthDate
	)

	_, err = db.Exec("UPDATE profile SET name = $1,linkedin = $2, description = $3, BirthDate = $4 WHERE id = $5", name, linkedin, description, BirthDate, id)
}

// Merges two profiles into one, Updates Second Profile with Firsts' Inputs
func MergeProfiles(prof1 *statics.Profile, prof2 *statics.Profile) {
	if prof1.Name != "" {
		prof2.Name = prof1.Name
	}
	if prof1.Description != "" {
		prof2.Description = prof1.Description
	}
	if prof1.LinkedIn != "" {
		prof2.LinkedIn = prof1.LinkedIn
	}
}
