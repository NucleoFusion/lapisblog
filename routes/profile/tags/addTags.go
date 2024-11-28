package tags

import (
	"database/sql"
	"net/http"
)

type addTagsRoute struct {
	Db *sql.DB
}

func (s *addTagsRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	key := r.URL.Query().Get("key")

}
