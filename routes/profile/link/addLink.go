package link

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type linkResponse struct {
	Id        int    `json:"id"`
	Key       string `json:"key"`
	LinkValue string `json:"linkValue"`
	LinkName  string `json:"linkName"`
}

type addLink struct {
	Db *sql.DB
}

func (s *addLink) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	r.ParseForm()
	body := r.PostForm

	valArr := body["linkValue"]
	nameArr := body["linkName"]
	if len(valArr) == 0 || len(nameArr) == 0 {
		io.WriteString(w, "invalid params")
		return
	}

	val := valArr[0]
	name := nameArr[0]

	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, "no key provided")
		return
	}

	userChan := make(chan int, 1)
	exists := make(chan bool, 1)

	InsertIntoLinks(name, val, key, s.Db, exists, userChan)

	if !(<-exists) {
		io.WriteString(w, "user not found")
		return
	}

	id := <-userChan

	resp := linkResponse{LinkValue: val, LinkName: name, Id: id, Key: key}

	data, _ := json.Marshal(resp)
	io.Writer.Write(w, data)
}

func InsertIntoLinks(linkName string, linkValue string, key string, db *sql.DB, exists chan bool, userChan chan int) {
	var id int

	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)
	if res.Scan(&id) == sql.ErrNoRows {
		exists <- false
		return
	}

	exists <- true
	userChan <- id

	_, err := db.Exec("INSERT INTO links(profile_id, link_name, link_value) VALUES ($1,$2,$3)", id, linkName, linkValue)
	if err != nil {
		fmt.Println(err)
	}
}
