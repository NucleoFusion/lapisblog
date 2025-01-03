package tags

import (
	"database/sql"
	"net/http"
)

func InitMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /add/{tag}", &addTagsRoute{Db: db})
	mux.Handle("GET /get", &GetTags{Db: db})
	mux.Handle("POST /remove/{id}", &RemoveTagRoute{Db: db})

	return mux
}
