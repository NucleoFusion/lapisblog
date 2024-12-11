package tags

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	lapiserror "lapisblog/lapisErrors"
	"net/http"
)

type GetTags struct {
	Db *sql.DB
}

type GetResponse struct {
	TagId   int    `json:"tagId"`
	TagName string `json:"tagName"`
}

func (s *GetTags) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Parsing Key and handling no Key provided
	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, lapiserror.NoKey)
		return
	}

	//Declaring Channels
	userExists := make(chan bool, 1)
	getResp := make(chan []GetResponse, 1)

	go GetFromDB(key, s.Db, userExists, getResp)

	if !(<-userExists) {
		io.WriteString(w, lapiserror.NoUser)
		return
	}

	resp := <-getResp

	data, _ := json.Marshal(resp)
	io.Writer.Write(w, data)
}

func GetFromDB(key string, db *sql.DB, userExists chan bool, getResp chan []GetResponse) {
	var id int
	resp := []GetResponse{}

	// Checking if valid key
	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)
	if res.Scan(&id) == sql.ErrNoRows {
		userExists <- false
		return
	}
	userExists <- true

	//Querying for all Tags associated with user
	result, err := db.Query("SELECT id, tag FROM tags WHERE user_id = $1", id)
	if err != nil {
		fmt.Println(err.Error())
		getResp <- resp
		return
	}

	for result.Next() {
		var (
			tagId   int
			tagName string
		)

		err = result.Scan(&tagId, &tagName)
		if err != nil {
			fmt.Println(err.Error())
			getResp <- resp
			return
		}

		strct := GetResponse{TagId: tagId, TagName: tagName}
		resp = append(resp, strct)
	}

	getResp <- resp
}
