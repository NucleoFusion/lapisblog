package profileRoute

import (
	"database/sql"
	"io"
	lapiserror "lapisblog/lapisErrors"
	"net/http"
	"strconv"
)

type ViewProfileRoute struct {
	Db *sql.DB
}

func (s *ViewProfileRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	_, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		io.WriteString(w, lapiserror.MissingParams)
		return
	}

}

func GetFromDB() {
}
