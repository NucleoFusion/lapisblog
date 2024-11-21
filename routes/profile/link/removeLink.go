package link

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type RemoveLinkRoute struct {
	Db *sql.DB
}

func (s *RemoveLinkRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	id, _ := strconv.Atoi(r.PathValue("id"))
	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, "key not found")
		return
	}

	userExists := make(chan bool, 1)
	linkExists := make(chan bool, 1)
	linkData := make(chan linkResponse, 1)

	RemoveLinkFromDB(s.Db, key, id, userExists, linkExists, linkData)

	if !(<-userExists) {
		io.WriteString(w, "invalid key, user not found")
		return
	}

	if !(<-linkExists) {
		io.WriteString(w, "no link with given id")
		return
	}

	link := <-linkData

	data, _ := json.Marshal(link)
	io.Writer.Write(w, data)
}

func RemoveLinkFromDB(db *sql.DB, key string, id int, userExists chan bool, linkExists chan bool, linkData chan linkResponse) {
	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)
	if res.Scan() == sql.ErrNoRows {
		userExists <- false
		return
	}
	userExists <- true

	resp := linkResponse{Key: key}

	var profID int

	linkRes := db.QueryRow("SELECT * FROM links WHERE id = $1", id)
	err := linkRes.Scan(&resp.Id, &profID, &resp.LinkName, &resp.LinkValue)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		linkExists <- false
		return
	}
	linkExists <- true

	linkData <- resp
	_, err2 := db.Exec("DELETE FROM links WHERE id = $1", id)
	if err2 != nil {
		fmt.Println(err)
	}
}
