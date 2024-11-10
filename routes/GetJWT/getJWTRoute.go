package getjwt

import (
	"database/sql"
	"fmt"
	"io"
	"lapisblog/auth"
	"net/http"
)

type GetJWT struct {
	Db *sql.DB
}

type JWTResponse struct {
	Jwt string `json:"jwt"`
}

func (s *GetJWT) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "plain/text")

	key := r.URL.Query().Get("key")
	fmt.Println("key: ", key)

	id, err := MatchKey(key, s.Db)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	payload := auth.CreatePayload(id, "User")
	jwt := payload.CreateJWT()
	jws, err := jwt.EncodeJWT()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.Writer.Write(w, []byte(jws))
}

func MatchKey(key string, db *sql.DB) (int, error) {
	var id int

	res := db.QueryRow("SELECT (id) FROM profile WHERE key = $1", key)

	err := res.Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
