package education

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	lapiserror "lapisblog/lapisErrors"
	"net/http"
	"strconv"
)

type RemoveEducationRoute struct {
	Db *sql.DB
}

type EducationDBResp struct {
	Id                   int     `json:"id"`
	ProfileId            int     `json:"profileID"`
	DegreeType           string  `json:"degreeType"`
	DegreeSpecialization string  `json:"degreeSpecialization"`
	GradeSystem          string  `json:"gradeSystem"`
	Grade                float64 `json:"grade"`
}

func (s *RemoveEducationRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	eduID, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		io.WriteString(w, lapiserror.InvalidPathParam)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, lapiserror.NoKey)
		return
	}

	edu := EducationDBResp{}

	userExists := make(chan bool, 1)
	eduFound := make(chan bool, 1)
	getEdu := make(chan EducationDBResp, 1)

	go RemoveFromDB(s.Db, key, eduID, &edu, userExists, eduFound, getEdu)

	//Handling case when user is not found
	if !(<-userExists) {
		io.WriteString(w, lapiserror.NoUser)
		return
	}

	//handling case when edu not found
	if !(<-eduFound) {
		io.WriteString(w, lapiserror.EduNotFound)
		return
	}

	data, _ := json.Marshal(<-getEdu)
	io.Writer.Write(w, data)
}

func RemoveFromDB(db *sql.DB, key string, eduId int64, edu *EducationDBResp, userExists chan bool, eduFound chan bool, getEdu chan EducationDBResp) {
	//Checking if the provided key actually corresponds to a user
	var id int

	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)
	if res.Scan(&id) == sql.ErrNoRows {
		userExists <- false
		return
	}
	userExists <- true

	//Deleting From DB
	res = db.QueryRow("DELETE FROM education WHERE id = $1 AND profile_id = $2 RETURNING *", eduId, id)
	err := res.Scan(&edu.Id, &edu.ProfileId, &edu.DegreeType, &edu.DegreeSpecialization, &edu.GradeSystem, &edu.Grade)
	if err == sql.ErrNoRows {
		eduFound <- false
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	eduFound <- true

	//Giving Data to Main Thread
	getEdu <- *edu
}
