package education

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	lapiserror "lapisblog/lapisErrors"
	"lapisblog/statics"
	"net/http"
)

type GetEducationRoute struct {
	Db *sql.DB
}

func (s *GetEducationRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, lapiserror.NoKey)
		return
	}

	userExists := make(chan bool, 1)
	sendData := make(chan *[]statics.Education)

	go GetFromDB(s.Db, key, userExists, sendData)

	if !(<-userExists) {
		io.WriteString(w, lapiserror.UserExists)
		return
	}

	arr := <-sendData

	data, _ := json.Marshal(arr)
	io.Writer.Write(w, data)
}

func GetFromDB(db *sql.DB, key string, userExists chan bool, sendData chan *[]statics.Education) {
	var id int

	//Checking if user exists
	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)
	if res.Scan(&id) == sql.ErrNoRows {
		userExists <- false
		return
	}
	userExists <- true

	//Getting data from education DB for given user
	result, err := db.Query("SELECT * FROM education WHERE profile_id = $1", id)
	if err != nil {
		fmt.Println(err.Error())
	}

	arr := ScanFromResult(result)
	sendData <- &arr
}

func ScanFromResult(result *sql.Rows) []statics.Education {
	arr := []statics.Education{}

	for result.Next() {
		var (
			id         int
			profile_id int
			edu        statics.Education
		)

		result.Scan(&id, &profile_id, &edu.DegreeType, &edu.DegreeSpecialization, &edu.GradeSystem, &edu.Grade)

		arr = append(arr, edu)
	}

	return arr
}
