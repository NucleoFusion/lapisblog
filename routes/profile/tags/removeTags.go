package tags

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	lapiserror "lapisblog/lapisErrors"
	"net/http"
	"strconv"
)

type RemoveTagRoute struct {
	Db *sql.DB
}

type RemoveTagResponse struct {
	Id  int    `json:"tagId"`
	Tag string `json:"tag"`
}

func (s *RemoveTagRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Parsing Key and handling no Key provided
	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, lapiserror.NoKey)
		return
	}

	// Getting ID of Tag to RemoveTagFromDB
	tag := r.PathValue("id")
	if tag == "" {
		io.WriteString(w, lapiserror.MissingParams)
		return
	}
	tag_id, _ := strconv.Atoi(tag)

	//Declaring Channels
	userExists := make(chan bool, 1)
	dbError := make(chan error, 1)
	getResp := make(chan *RemoveTagResponse, 1)

	go RemoveTagFromDB(key, tag_id, s.Db, userExists, dbError, getResp)

	if !(<-userExists) {
		io.WriteString(w, lapiserror.NoUser)
		return
	}

	err := <-dbError
	if err != nil {
		io.WriteString(w, fmt.Sprintf("ERROR: %s", err.Error()))
		return
	}

	data, _ := json.Marshal(*(<-getResp))
	io.Writer.Write(w, data)
}

func RemoveTagFromDB(key string, tag_id int, db *sql.DB, userExists chan bool, dbError chan error, getResp chan *RemoveTagResponse) {
	var id int

	//Querying Database to check whether user exists
	userRes := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)

	err := userRes.Scan(&id)
	if err == sql.ErrNoRows {
		userExists <- false
	} else if err != nil {
		fmt.Println(err.Error())
	}
	userExists <- true

	// Deleting From Database
	resp := RemoveTagResponse{}

	tagRes := db.QueryRow("DELETE FROM tags WHERE id = $1 AND user_id = $2 RETURNING id, tag", tag_id, id)
	err = tagRes.Scan(&resp.Id, &resp.Tag)

	//Sending to Main thread if any error occurred
	dbError <- err

	// Sending to user the response
	getResp <- &resp
}
