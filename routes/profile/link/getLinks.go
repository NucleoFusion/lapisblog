package link

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LinkDBResponse struct {
	Id    uint8  `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetLinkRoute struct {
	Db *sql.DB
}

func (s *GetLinkRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	key := r.URL.Query().Get("key")
	if key == "" {
		io.WriteString(w, "key not found")
		return
	}

	exists := make(chan bool, 1)
	dataChan := make(chan *[]LinkDBResponse, 1)

	go GetFromDB(s.Db, key, exists, dataChan)

	if !(<-exists) {
		io.WriteString(w, "no user with given key")
		return
	}

	data := <-dataChan

	byteData, _ := json.Marshal(data)
	io.Writer.Write(w, byteData)
}

func GetFromDB(db *sql.DB, key string, exists chan bool, dataChan chan *[]LinkDBResponse) {
	var id int

	res := db.QueryRow("SELECT (id) FROM profile where key = $1", key)
	if res.Scan(&id) == sql.ErrNoRows {
		exists <- false
	}
	exists <- true

	result, err := db.Query("SELECT * FROM links WHERE profile_id = $1", id)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer result.Close()

	var arr []LinkDBResponse

	for result.Next() {
		curr := LinkDBResponse{}

		var profileID int

		err := result.Scan(&curr.Id, &profileID, &curr.Name, &curr.Value)
		if err != nil {
			fmt.Println(err.Error())
		}

		arr = append(arr, curr)
	}

	dataChan <- &arr
}
