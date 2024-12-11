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
	UserId int      `json:"userId"`
	Tags   []string `json:"tags"`
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
	tagExists := make(chan bool, 1)
	getResp := make(chan *TagResponse, 1)

	go InsertIntoDB(s.Db, key, tag, userExists, tagExists, getResp)

	if !(<-userExists) {
		io.WriteString(w, lapiserror.NoUser)
		return
	}

	if <-tagExists {
		io.WriteString(w, lapiserror.TagExists)
		return
	}

	// Getting Response Struct
	resp := <-getResp

	data, _ := json.Marshal(resp)
	io.Writer.Write(w, data)
}

func InsertIntoDB(db *sql.DB, key string, tag enums.Tags, userExists chan bool, tagExists chan bool, getResp chan *TagResponse) {
	var id int
	resp := TagResponse{UserId: id, Tags: []string{}}

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

	//Getting Existing Tags from DB
	result, err := db.Query("SELECT (tag) FROM tags WHERE user_id = $1", id)
	if err != nil {
		fmt.Println(err.Error())
		tagExists <- false
		getResp <- &resp
		return
	}

	for result.Next() {
		var tagName string

		err := result.Scan(&tagName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//Checking whether this tag is the same as the one we have to add
		if tagName == string(tag) {
			tagExists <- true
			return
		}

		resp.Tags = append(resp.Tags, tagName)
	}
	tagExists <- false

	//Adding the Current Tag
	resp.Tags = append(resp.Tags, string(tag))
	getResp <- &resp

	//Adding Tag to DB
	db.Exec("INSERT INTO tags(user_id, tag) VALUES ($1,$2) RETURNING *", id, tag)
}
