package education

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	lapiserror "lapisblog/lapisErrors"
	"lapisblog/statics"
	enums "lapisblog/statics/Enums"
	"net/http"
	"net/url"
	"strconv"
)

type AddEducation struct {
	Db *sql.DB
}

func (s *AddEducation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	r.ParseForm()
	body := r.PostForm

	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, lapiserror.NoKey)
		return
	}

	edu := statics.Education{}

	//Declaring Decode Chans
	decodeError := make(chan error, 1)
	decodeDone := make(chan bool, 1)

	//Declaring Insert Chans
	userExists := make(chan bool, 1)
	decodeDoneInsert := make(chan bool, 1)

	go DecodeBody(&edu, &body, decodeError, decodeDone)
	go InsertIntoDB(s.Db, &edu, key, userExists, decodeDoneInsert)

	// Seeing if Decode was successful or if an error occurred
	err := <-decodeError
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	//Waiting for Decode to Complete, Then Starting the 2nd part of InsertIntoDB()
	<-decodeDone

	//Telling InsertIntoDB to start inserting
	decodeDoneInsert <- true

	if !(<-userExists) {
		io.WriteString(w, lapiserror.NoUser)
		return
	}

	data, _ := json.Marshal(&edu)
	io.Writer.Write(w, data)
}

func InsertIntoDB(db *sql.DB, edu *statics.Education, key string, userExists chan bool, decodeDoneInsert chan bool) {
	var id int

	//Scanning for whether key is valid
	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)
	if res.Scan(&id) == sql.ErrNoRows {
		userExists <- false
		return
	}

	userExists <- true

	//Waiting for Decode to Complete
	<-decodeDoneInsert

	_, err := db.Exec("INSERT INTO education(profile_id, degree_type, degree_specialization, grade_system, grade) VALUES ( $1, $2, $3, $4, $5)", id, edu.DegreeType, edu.DegreeSpecialization, edu.GradeSystem, edu.Grade)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DecodeBody(edu *statics.Education, body *url.Values, decodeError chan error, decodeDone chan bool) {
	for key, val := range *body {
		switch key {
		case "degreeType":
			typ, err := enums.GetDegree(val[0])
			if err != nil {
				decodeError <- err
				decodeDone <- true
				return
			}

			edu.DegreeType = typ

		case "gradeSystem":
			typ, err := enums.GetGradeSys(val[0])
			if err != nil {
				decodeError <- err
				decodeDone <- true
				return
			}

			edu.GradeSystem = typ

		case "grade":
			i, _ := strconv.ParseFloat(val[0], 32)
			edu.Grade = float32(i)

		case "specialization":
			edu.DegreeSpecialization = val[0]

		}
	}

	if edu.DegreeSpecialization == "" || edu.Grade == float32(0) || edu.DegreeType == "" || edu.GradeSystem == "" {
		decodeError <- errors.New(lapiserror.MissingParams)
		decodeDone <- true
		return
	}

	decodeError <- nil
	decodeDone <- true
}
