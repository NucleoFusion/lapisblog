package tags

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	lapiserror "lapisblog/lapisErrors"
	enums "lapisblog/statics/Enums"
	"net/http"
)

type addTagsRoute struct {
	Db *sql.DB
}

type TagResponse struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Tag    string `json:"tag"`
}

func (s *addTagsRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// Parsing Key and handling no Key provided
	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, lapiserror.NoKey)
		return
	}

	// Getting Tag value and checking if Valid Tag
	tag, err := enums.GetTag(r.PathValue("tag"))
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	// Declaring channel to communicate with InsertIntoDB Goroutine
	userExists := make(chan bool, 1)
	getResp := make(chan *TagResponse, 1)

	InsertIntoDB(s.Db, key, tag, userExists, getResp)

	if !(<-userExists) {
		io.WriteString(w, lapiserror.NoUser)
		return
	}

	// Getting Response Struct
	resp := <-getResp

	data, _ := json.Marshal(resp)
	io.Writer.Write(w, data)
}

func InsertIntoDB(db *sql.DB, key string, tag enums.Tags, userExists chan bool, getResp chan *TagResponse) {
	var id int
	resp := TagResponse{}

	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)

	//Checking whether user with given key actually exists
	err := res.Scan(&id)
	if err == sql.ErrNoRows {
		userExists <- false
		return
	} else if err != nil {
		fmt.Println(err.Error())
		return
	}
	userExists <- true

	//Adding Tag to DB
	result := db.QueryRow("INSERT INTO tags(user_id, tag) VALUES ($1,$2) RETURNING *", id, tag)
	result.Scan(&resp.Id, &resp.UserId, &resp.Tag)

	getResp <- &resp
}
